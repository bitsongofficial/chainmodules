package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateAuction checks if the given auction is valid
func ValidateAuction(auction Auction) error {
	if len(auction.Owner) > 0 {
		if _, err := sdk.AccAddressFromBech32(auction.Owner); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
		}
	}
	if err := ValidateAuctionType(auction.AuctionType); err != nil {
		return err
	}
	if err := ValidateDuration(auction.Duration); err != nil {
		return err
	}
	if err := ValidateAmount(auction.MinAmount.Amount); err != nil {
		return err
	}
	if err := ValidateAuctionLimit(auction.AuctionType, auction.Limit); err != nil {
		return err
	}
	return nil
}

// ValidateAuctionType checks if the given auction type is valid
func ValidateAuctionType(auctionType uint32) error {
	if auctionType > 2 {
		return sdkerrors.Wrapf(ErrInvalidAuctionType, "invalid auction type (%d)", auctionType)
	}
	return nil
}

// ValidateAuctionLimit checks if the given auction limit is valid
func ValidateAuctionLimit(auctionType uint32, limit uint32) error {
	if (auctionType == 0 && limit != 1) || (auctionType == 1 && limit != 0) {
		return sdkerrors.Wrapf(ErrInvalidAuctionLimit, "invalid auction limit (%d)", limit)
	}
	return nil
}

// ValidateAmount checks if the given amount is valid
func ValidateAmount(amount sdk.Int) error {
	if amount.IsZero() {
		return sdkerrors.Wrapf(ErrInvalidAmount, "invalid amount %d, only accepts positive amount", amount)
	}
	return nil
}

// ValidateDuration checks if the given duration is valid
func ValidateDuration(duration uint64) error {
	if duration == 0 {
		return sdkerrors.Wrapf(ErrInvalidDuration, "invalid duration %d, only accepts positive value", duration)
	}
	return nil
}

// ValidateBid checks if the given bid is valid
func ValidateBid(bid Bid) error {
	if len(bid.Bidder) > 0 {
		if _, err := sdk.AccAddressFromBech32(bid.Bidder); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bidder address (%s)", err)
		}
	}
	if err := ValidateAmount(bid.BidAmount.Amount); err != nil {
		return err
	}
	return nil
}
