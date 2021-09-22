package types

import (
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ proto.Message = &Auction{}
)

// AuctionI defines an interface for Auction
type AuctionI interface {
	GetId() uint64
	GetAuctionType() uint32
	GetNftId() string
	GetStartTime() uint64
	GetDuration() uint64
	GetMinAmount() sdk.Coin
	GetOwner() sdk.AccAddress
	GetLimit() uint32
}

// NewAuction constructs a new Auction instance
func NewAuction(
	id uint64,
	auctionType uint32,
	nftId string,
	startTime uint64,
	duration uint64,
	minAmount sdk.Coin,
	owner sdk.AccAddress,
	limit uint32,
) Auction {
	return Auction{
		Id:          id,
		AuctionType: auctionType,
		NftId:       nftId,
		StartTime:   startTime,
		Duration:    duration,
		MinAmount:   minAmount,
		Owner:       owner.String(),
		Limit:       limit,
	}
}

// GetId implements exported.AuctionI
func (t Auction) GetId() uint64 {
	return t.Id
}

// GetAuctionType implements exported.AuctionI
func (t Auction) GetAuctionType() uint32 {
	return t.AuctionType
}

// GetNftId implements exported.AuctionI
func (t Auction) GetNftId() string {
	return t.NftId
}

// GetStartTime implements exported.AuctionI
func (t Auction) GetStartTime() uint64 {
	return t.StartTime
}

// GetDuration implements exported.AuctionI
func (t Auction) GetDuration() uint64 {
	return t.Duration
}

// GetMinAmount implements exported.AuctionI
func (t Auction) GetMinAmount() sdk.Coin {
	return t.MinAmount
}

// GetOwner implements exported.AuctionI
func (t Auction) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(t.Owner)
	return owner
}

// GetLimit implements exported.AuctionI
func (t Auction) GetLimit() uint32 {
	return t.Limit
}
