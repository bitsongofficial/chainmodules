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
	storeKey      sdk.StoreKey
	cdc           codec.Marshaler
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	nftKeeper     types.NftKeeper
}

func NewKeeper(
	cdc codec.Marshaler,
	key sdk.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	nftKeeper types.NftKeeper,
) Keeper {
	return Keeper{
		storeKey:   key,
		cdc:        cdc,
		bankKeeper: bankKeeper,
		nftKeeper:  nftKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("go-bitsong/%s", types.ModuleName))
}

// OpenAuction opens a new auction
func (k Keeper) OpenAuction(
	ctx sdk.Context,
	auctionType types.AuctionType,
	nftDenomId string,
	nftTokenId string,
	duration uint64,
	minAmount sdk.Coin,
	owner sdk.AccAddress,
	limit uint32,
) (uint64, error) {
	id := k.GetLastAuctionId(ctx)
	now := time.Now()
	auction := types.NewAuction(id, auctionType, nftDenomId, nftTokenId, uint64(now.Unix()), duration, minAmount, owner, limit)

	if err := k.AddAuction(ctx, auction); err != nil {
		return 0, err
	}
	k.SetLastAuctionId(ctx, id+1)

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

	k.cancelAuction(ctx, auction)

	return nil
}

// OpenBid opens a new bid
func (k Keeper) OpenBid(
	ctx sdk.Context,
	auctionId uint64,
	bidder sdk.AccAddress,
	bidAmount sdk.Coin,
) error {
	auction, err := k.getAuctionById(ctx, auctionId)
	if err != nil {
		return err
	}

	if auction.GetStatus() != types.AUCTION_STATUS_RUNNING {
		return sdkerrors.Wrapf(types.ErrInvalidAuction, "the auction was already ended or cancelled")
	}

	if bidAmount.Denom != auction.GetMinAmount().Denom {
		return sdkerrors.Wrapf(types.ErrInvalidBidAmountDenom, "bid amount denom is different with auction min amount denom")
	}

	if bidAmount.Amount.LT(auction.GetMinAmount().Amount) {
		return sdkerrors.Wrapf(types.ErrNotEnoughBidAmount, "the bid amount is less than the auction min amount")
	}

	bids := k.getBidsByAuctionId(ctx, auctionId)
	if auction.GetLimit() != 0 {
		if len(bids) >= int(auction.GetLimit()) {
			var minBid types.Bid = bids[0]
			for _, bid := range bids {
				if bid.GetBidAmount().Amount.LT(minBid.GetBidAmount().Amount) {
					minBid = bid
				}
			}
			if bidAmount.Amount.LTE(minBid.GetBidAmount().Amount) {
				return sdkerrors.Wrapf(types.ErrNotEnoughBidAmount, "the bid amount is not enough to be a winner")
			} else {
				err = k.cancelBid(ctx, auctionId, minBid.GetBidder())
				if err != nil {
					return err
				}
			}
		}
	}

	bid := types.NewBid(auctionId, bidder, bidAmount)

	if err := k.AddBid(ctx, bid); err != nil {
		return err
	}

	return nil
}

// CancelBid cancels the specified bid
func (k Keeper) CancelBid(
	ctx sdk.Context,
	auctionId uint64,
	bidder sdk.AccAddress,
) error {
	if err := k.cancelBid(ctx, auctionId, bidder); err != nil {
		return err
	}

	return nil
}

// Withdraw sends nft to the auction winner
func (k Keeper) Withdraw(
	ctx sdk.Context,
	auctionId uint64,
	recipient sdk.AccAddress,
) error {
	auction, err := k.getAuctionById(ctx, auctionId)
	if err != nil {
		return err
	}

	if recipient.Equals(auction.GetOwner()) {
		err = k.withdrawCoins(ctx, auction, recipient)
	} else {
		err = k.withdrawNFT(ctx, auction, recipient)
	}

	if err != nil {
		return err
	}

	return nil
}
