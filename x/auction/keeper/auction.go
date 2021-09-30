package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
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

		auction, err := k.getAuctionById(ctx, id.Value)
		if err != nil {
			continue
		}
		auctions = append(auctions, auction)
	}
	return
}

func (k Keeper) getAuctionById(ctx sdk.Context, id uint64) (auction types.Auction, err error) {
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
	if k.hasAuction(ctx, auction.GetId()) {
		return sdkerrors.Wrapf(types.ErrAuctionAlreadyExists, "auction already exists: %d", auction.GetId())
	}

	// set auction
	k.setAuction(ctx, auction)

	if len(auction.GetOwner()) != 0 {
		// set token to be prefixed with owner
		k.setWithOwner(ctx, auction.GetOwner(), auction.GetId())
	}

	return nil
}

func (k Keeper) hasAuction(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyAuctionById(id))
}

func (k Keeper) setAuction(ctx sdk.Context, auction types.Auction) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&auction)

	store.Set(types.KeyAuctionById(auction.GetId()), bz)
}

func (k Keeper) setWithOwner(ctx sdk.Context, owner sdk.AccAddress, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&gogotypes.UInt64Value{Value: id})

	store.Set(types.KeyAuctionsByOwner(owner, id), bz)
}

// AddAuction saves a new auction
func (k Keeper) withdraw(ctx sdk.Context, auctionId uint64, recipient sdk.AccAddress) error {
	if !k.hasAuction(ctx, auctionId) {
		return sdkerrors.Wrapf(types.ErrAuctionNotExists, "auction not found: %d", auctionId)
	}

	return nil
}
