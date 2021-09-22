package types

import (
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ proto.Message = &Bid{}
)

// BidI defines an interface for Bid
type BidI interface {
	GetAuctionId() uint64
	GetBidder() string
	GetBidAmount() sdk.Coin
}

// NewBid constructs a new Bid instance
func NewBid(
	auctionId uint64,
	bidder sdk.AccAddress,
	bidAmount sdk.Coin,
) Bid {
	return Bid{
		AuctionId: auctionId,
		Bidder:    bidder.String(),
		BidAmount: bidAmount,
	}
}

// GetAuctionId implements exported.BidI
func (t Bid) GetAuctionId() uint64 {
	return t.AuctionId
}

// GetBidder implements exported.BidI
func (t Bid) GetBidder() sdk.AccAddress {
	bidder, _ := sdk.AccAddressFromBech32(t.Bidder)
	return bidder
}

// GetBidAmount implements exported.BidI
func (t Bid) GetBidAmount() sdk.Coin {
	return t.BidAmount
}
