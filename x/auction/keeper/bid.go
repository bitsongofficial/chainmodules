package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

func (k Keeper) getBid(ctx sdk.Context, auctionId uint64, bidder sdk.AccAddress) (bid types.Bid, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyBid(auctionId, bidder))
	if bz == nil {
		return bid, sdkerrors.Wrap(types.ErrBidNotExists, fmt.Sprintf("bid %s does not exist", bidder))
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &bid)
	return bid, nil
}

func (k Keeper) GetAllBids(ctx sdk.Context) (bids []types.Bid) {
	store := ctx.KVStore(k.storeKey)

	var it sdk.Iterator = sdk.KVStorePrefixIterator(store, types.PrefixBidsByAuctionId)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var bid types.Bid
		k.cdc.MustUnmarshalBinaryBare(it.Value(), &bid)

		bids = append(bids, bid)
	}
	return
}

func (k Keeper) getBidsByAuctionId(ctx sdk.Context, auctionId uint64) (bids []types.Bid) {
	store := ctx.KVStore(k.storeKey)

	var it sdk.Iterator = sdk.KVStorePrefixIterator(store, types.KeyBidsByAuctionId(auctionId))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var bid types.Bid
		k.cdc.MustUnmarshalBinaryBare(it.Value(), &bid)

		bids = append(bids, bid)
	}
	return
}

func (k Keeper) getBidsByBidder(ctx sdk.Context, bidder sdk.AccAddress) (bids []types.Bid) {
	store := ctx.KVStore(k.storeKey)

	var it sdk.Iterator = sdk.KVStorePrefixIterator(store, append(types.PrefixBidsByBidder, bidder.Bytes()...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var id gogotypes.UInt64Value
		k.cdc.MustUnmarshalBinaryBare(it.Value(), &id)

		bid, err := k.getBid(ctx, id.Value, bidder)
		if err != nil {
			continue
		}

		bids = append(bids, bid)
	}
	return
}

// AddBid saves a new bid
func (k Keeper) AddBid(ctx sdk.Context, bid types.Bid) error {
	if k.hasBid(ctx, bid.GetAuctionId(), bid.GetBidder()) {
		return sdkerrors.Wrapf(types.ErrBidAlreadyExists, "bid already exists: %s", bid.GetBidder())
	}

	// set bid
	k.setBid(ctx, bid)

	if len(bid.GetBidder()) != 0 {
		k.setWithBidder(ctx, bid.GetBidder(), bid.GetAuctionId())
	}

	return nil
}

// CancelBid removes a bid
func (k Keeper) cancelBid(ctx sdk.Context, auctionId uint64, bidder sdk.AccAddress) error {
	if !k.hasBid(ctx, auctionId, bidder) {
		return sdkerrors.Wrapf(types.ErrBidNotExists, "bid not found: %s", bidder)
	}

	k.removeBid(ctx, auctionId, bidder)

	if len(bidder) != 0 {
		k.removeWithBidder(ctx, bidder, auctionId)
	}

	return nil
}

func (k Keeper) hasBid(ctx sdk.Context, auctionId uint64, bidder sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyBid(auctionId, bidder))
}

func (k Keeper) setBid(ctx sdk.Context, bid types.Bid) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&bid)

	store.Set(types.KeyBid(bid.GetAuctionId(), bid.GetBidder()), bz)
}

func (k Keeper) setWithBidder(ctx sdk.Context, bidder sdk.AccAddress, auctionId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&gogotypes.UInt64Value{Value: auctionId})

	store.Set(types.KeyBidsByBidder(bidder, auctionId), bz)
}

func (k Keeper) removeBid(ctx sdk.Context, auctionId uint64, bidder sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.KeyBid(auctionId, bidder))
}

func (k Keeper) removeWithBidder(ctx sdk.Context, bidder sdk.AccAddress, auctionId uint64) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.KeyBidsByBidder(bidder, auctionId))
}
