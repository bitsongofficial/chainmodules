package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.Marshaler
	bankKeeper types.BankKeeper
}

func NewKeeper(
	cdc codec.Marshaler,
	key sdk.StoreKey,
	bankKeeper types.BankKeeper,
) Keeper {
	return Keeper{
		storeKey:   key,
		cdc:        cdc,
		bankKeeper: bankKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("go-bitsong/%s", types.ModuleName))
}

// OpenAuction opens a new auction
func (k Keeper) OpenAuction(
	ctx sdk.Context,
	auctionType uint32,
	nftId string,
	duration uint64,
	minAmount sdk.Coin,
	owner sdk.AccAddress,
	limit uint32,
) (uint64, error) {
	id := k.getLastAuctionId(ctx)
	now := time.Now()
	auction := types.NewAuction(id, auctionType, nftId, uint64(now.Unix()), duration, minAmount, owner, limit)

	if err := k.AddAuction(ctx, auction); err != nil {
		return err
	}
	k.setLastAuctionId(ctx, id+1)

	return id, nil
}

// EditAuction edits the specified auction
func (k Keeper) EditAuction(
	ctx sdk.Context,
	id uint64,
	duration uint64,
	owner sdk.AccAddress,
) error {
	// get the destination auction
	auction, err := k.getAuctionById(ctx, id)
	if err != nil {
		return err
	}

	if owner.String() != auction.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the auction %d", owner, id)
	}

	if auction.GetStatus() != types.AUCTION_STATUS_RUNNING {
		return sdkerrors.Wrapf(types.ErrInvalidAuction, "the auction was already ended or cancelled")
	}

	now := time.Now()
	if auction.StartTime+duration < uint64(now.Unix()) {
		return sdkerrors.Wrapf(types.ErrInvalidDuration, "startTime + duration < now")
	}

	auction.Duration = duration

	k.setAuction(ctx, auction)

	return nil
}

// CancelAuction cancels the specified auction
func (k Keeper) CancelAuction(
	ctx sdk.Context,
	id uint64,
	owner sdk.AccAddress,
) error {
	// get the destination auction
	auction, err := k.getAuctionById(ctx, id)
	if err != nil {
		return err
	}

	if owner.String() != auction.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the auction %d", owner, id)
	}

	if auction.GetStatus() != types.AUCTION_STATUS_RUNNING {
		return sdkerrors.Wrapf(types.ErrInvalidAuction, "the auction was already ended or cancelled")
	}

	auction.Cancelled = true

	k.setAuction(ctx, auction)

	return nil
}
