package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
	nfttypes "github.com/bitsongofficial/chainmodules/x/nft/types"
)

func (k Keeper) GetLastAuctionId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.KeyLastAuctionId)
	if bytes == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bytes)
}

func (k Keeper) SetLastAuctionId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLastAuctionId, sdk.Uint64ToBigEndian(id))
}

func (k Keeper) GetAuctions(ctx sdk.Context, owner sdk.AccAddress) (auctions []types.Auction) {
	store := ctx.KVStore(k.storeKey)

	var it sdk.Iterator
	if owner == nil {
		it = sdk.KVStorePrefixIterator(store, types.PrefixAuctionById)
		defer it.Close()

		for ; it.Valid(); it.Next() {
			var auction types.Auction
			k.cdc.MustUnmarshalBinaryBare(it.Value(), &auction)

			auctions = append(auctions, auction)
		}
		return
	}

	it = sdk.KVStorePrefixIterator(store, append(types.PrefixAuctionsByOwner, owner.Bytes()...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var id gogotypes.UInt64Value
		k.cdc.MustUnmarshalBinaryBare(it.Value(), &id)

		auction, err := k.GetAuctionById(ctx, id.Value)
		if err != nil {
			continue
		}
		auctions = append(auctions, auction)
	}
	return
}

func (k Keeper) GetAuctionById(ctx sdk.Context, id uint64) (auction types.Auction, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyAuctionById(id))
	if bz == nil {
		return auction, sdkerrors.Wrap(types.ErrAuctionNotExists, fmt.Sprintf("auction %d does not exist", id))
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &auction)
	return auction, nil
}

// AddAuction saves a new auction
func (k Keeper) AddAuction(ctx sdk.Context, auction types.Auction) error {
	if k.HasAuction(ctx, auction.GetId()) {
		return sdkerrors.Wrapf(types.ErrAuctionAlreadyExists, "auction already exists: %d", auction.GetId())
	}

	// set auction
	k.SetAuction(ctx, auction)

	if len(auction.GetOwner()) != 0 {
		// set token to be prefixed with owner
		k.setWithOwner(ctx, auction.GetOwner(), auction.GetId())
	}

	err := k.nftKeeper.TransferOwner(ctx, auction.GetNftDenomId(), auction.GetNftTokenId(), auction.GetOwner(), k.accountKeeper.GetModuleAddress(types.ModuleName))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) cancelAuction(ctx sdk.Context, auction types.Auction) error {
	auction.Cancelled = true
	k.SetAuction(ctx, auction)

	bids := k.GetBidsByAuctionId(ctx, auction.GetId())
	for _, bid := range bids {
		if err := k.cancelBid(ctx, auction.GetId(), bid.GetBidder()); err != nil {
			return err
		}
	}

	err := k.nftKeeper.TransferOwner(ctx, auction.NftDenomId, auction.NftTokenId, k.accountKeeper.GetModuleAddress(types.ModuleName), auction.GetOwner())
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) HasAuction(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyAuctionById(id))
}

func (k Keeper) SetAuction(ctx sdk.Context, auction types.Auction) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&auction)

	store.Set(types.KeyAuctionById(auction.GetId()), bz)
}

func (k Keeper) setWithOwner(ctx sdk.Context, owner sdk.AccAddress, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&gogotypes.UInt64Value{Value: id})

	store.Set(types.KeyAuctionsByOwner(owner, id), bz)
}

// withdrawCoins send auction coins to the recipient(splitshare, royaltyshare applies here)
func (k Keeper) withdrawCoins(ctx sdk.Context, auction types.Auction, recipient sdk.AccAddress) error {
	bids := k.GetBidsByAuctionId(ctx, auction.GetId())
	amount := sdk.NewCoin(bids[0].GetBidAmount().Denom, sdk.ZeroInt())
	for _, bid := range bids {
		amount.Amount = amount.Amount.Add(bid.GetBidAmount().Amount)
	}

	denom, err := k.nftKeeper.GetDenom(ctx, auction.GetNftDenomId())
	if err != nil {
		return err
	}

	item, err := k.nftKeeper.GetNFT(ctx, auction.GetNftDenomId(), auction.GetNftTokenId())
	if err != nil {
		return err
	}

	multiplier := sdk.NewInt(100)
	if !item.GetPrimaryStatus() {
		multiplier = denom.RoyaltyShare.RoundInt()
	}

	recipientAmount := amount.Amount

	for i := 0; i < len(denom.Creators); i++ {
		creatorAmount := amount.Amount.Mul(denom.SplitShares[i].RoundInt()).Mul(multiplier).Quo(sdk.NewInt(10000))
		creatorAddr, err := sdk.AccAddressFromBech32(denom.Creators[i])
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddr, sdk.Coins{sdk.NewCoin(amount.Denom, creatorAmount)})
		if err != nil {
			return err
		}
		recipientAmount = recipientAmount.Sub(creatorAmount)
	}

	if !item.GetPrimaryStatus() {
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, sdk.Coins{sdk.NewCoin(amount.Denom, recipientAmount)})
	} else {
		k.nftKeeper.SetNFT(ctx, auction.GetNftDenomId(), nfttypes.NewBaseNFT(item.GetID(), item.GetName(), item.GetOwner(), item.GetURI(), false))
	}

	if auction.GetAuctionType() != types.Single_Edition {
		err := k.nftKeeper.TransferOwner(ctx, auction.NftDenomId, auction.NftTokenId, k.accountKeeper.GetModuleAddress(types.ModuleName), auction.GetOwner())
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) GetAuctionStatus(ctx sdk.Context, auction types.Auction) types.AuctionStatus {
	if auction.Cancelled {
		return types.AUCTION_STATUS_CANCELLED
	}
	now := ctx.BlockTime()
	if auction.StartTime+auction.Duration < uint64(now.Unix()) {
		return types.AUCTION_STATUS_ENDED
	}
	return types.AUCTION_STATUS_RUNNING
}
