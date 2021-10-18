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
	if err := ValidateAmount(auction.AuctionType, auction.MinAmount.Amount); err != nil {
		return err
	}
	if err := ValidateAuctionLimit(auction.AuctionType, auction.Limit); err != nil {
		return err
	}
	return nil
}

// ValidateAuctionType checks if the given auction type is valid
func ValidateAuctionType(auctionType AuctionType) error {
	if auctionType > 2 {
		return sdkerrors.Wrapf(ErrInvalidAuctionType, "invalid auction type (%d)", auctionType)
	}
	return nil
}

// ValidateAuctionLimit checks if the given auction limit is valid
func ValidateAuctionLimit(auctionType AuctionType, limit uint32) error {
	if (auctionType == Single_Edition && limit != 1) || (auctionType == Open_Edition && limit != 0) || (auctionType != Open_Edition && limit == 0) {
		return sdkerrors.Wrapf(ErrInvalidAuctionLimit, "invalid auction limit (%d)", limit)
	}
	return nil
}

// ValidateAmount checks if the given amount is valid
func ValidateAmount(auctionType AuctionType, amount sdk.Int) error {
	if auctionType != Open_Edition && amount.IsZero() {
		return sdkerrors.Wrapf(ErrInvalidAmount, "invalid amount %d, only accepts positive amount", amount)
	}
	if auctionType == Open_Edition && !amount.IsZero() {
		return sdkerrors.Wrapf(ErrInvalidAmount, "invalid amount %d, only accepts zero amount", amount)
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
	} else {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "nil bidder address")
	}
	return nil
}
