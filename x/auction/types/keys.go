package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the auction module
	ModuleName = "auction"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the auction module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the auction module
	RouterKey string = ModuleName

	// DefaultParamspace is the default name for parameter store
	DefaultParamspace = ModuleName
)

var (
	// PrefixAuctionById defines a prefix for the auction by id
	PrefixAuctionById = []byte{0x01}
	// PrefixAuctionsByOwner defines a prefix for the auctions by owner
	PrefixAuctionsByOwner = []byte{0x02}
	// PrefixBidsByAuctionId defines a prefix for the bids by auction id
	PrefixBidsByAuctionId = []byte{0x03}
	// PrefixBidsByOwner defines a prefix for the bids by owner
	PrefixBidsByOwner = []byte{0x04}
	// KeyLastAuctionId defines the key for the last auction id
	KeyLastAuctionId = []byte{0x05}
)

// KeyAuctionById returns the key of the specified id. Intended for querying the auction by the id.
func KeyAuctionById(id uint64) []byte {
	return append(PrefixAuctionById, sdk.Uint64ToBigEndian(id)...)
}

// KeyAuctionsByOwner returns the key of the specified owner. Intended for querying all auctions of an owner
func KeyAuctionsByOwner(owner sdk.AccAddress, id uint64) []byte {
	return append(append(PrefixAuctionsByOwner, owner.Bytes()...), sdk.Uint64ToBigEndian(id)...)
}

// KeyBid returns the key of the specified auction id and bidder. Intended for querying a bid of an auction
func KeyBid(id uint64, bidder sdk.AccAddress) []byte {
	return append(append(PrefixBidsByAuctionId, sdk.Uint64ToBigEndian(id)...), bidder.Bytes()...)
}

// KeyBidsByAuctionId returns the key of the specified auction id. Intended for querying all bids of an auction
func KeyBidsByAuctionId(id uint64) []byte {
	return append(PrefixBidsByAuctionId, sdk.Uint64ToBigEndian(id)...)
}

// KeyBidsByOwner returns the key of the specified owner. Intended for querying all bids of an owner
func KeyBidsByBidder(bidder sdk.AccAddress, auctionId uint64) []byte {
	return append(append(PrefixBidsByOwner, bidder.Bytes()...), sdk.Uint64ToBigEndian(auctionId)...)
}
