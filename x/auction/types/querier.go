package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QueryAuction         = "auction"
	QueryAllAuctions     = "auctions"
	QueryAuctionsByOwner = "auctions-by-owner"
	QueryBid             = "bid"
	QueryBids            = "bids"
	QueryBidsByBidder    = "bids-by-bidder"
)

// QueryAuctionParams is the query parameters for 'custom/auction/auction/{id}'
type QueryAuctionParams struct {
	Id uint64
}

// NewQueryAuctionParams creates a new instance of QueryAuctionParams
func NewQueryAuctionParams(id uint64) QueryAuctionParams {
	return QueryAuctionParams{
		Id: id,
	}
}

// QueryAllAuctionsParams is the query parameters for 'custom/auction/auctions'
type QueryAllAuctionsParams struct {
}

// NewQueryAllAuctionsParams creates a new instance of QueryAllAuctionsParams
func NewQueryAllAuctionsParams() QueryAllAuctionsParams {
	return QueryAllAuctionsParams{}
}

// QueryAuctionsByOwnerParams is the query parameters for 'custom/auction/auctions/{owner}'
type QueryAuctionsByOwnerParams struct {
	Owner sdk.AccAddress
}

// NewQueryAuctionsByOwnerParams creates a new instance of QueryAuctionsByOwnerParams
func NewQueryAuctionsByOwnerParams(owner sdk.AccAddress) QueryAuctionsByOwnerParams {
	return QueryAuctionsByOwnerParams{
		Owner: owner,
	}
}

// QueryBidsParams is the query parameters for 'custom/auction/bids'
type QueryBidParams struct {
	AuctionId uint64
	Bidder    sdk.AccAddress
}

// NewQueryBidParams creates a new instance of QueryBidParams
func NewQueryBidParams(auctionId uint64, bidder sdk.AccAddress) QueryBidParams {
	return QueryBidParams{
		AuctionId: auctionId,
		Bidder:    bidder,
	}
}

// QueryBidsParams is the query parameters for 'custom/auction/bids'
type QueryBidsParams struct {
	AuctionId uint64
}

// NewQueryBidsParams creates a new instance of QueryBidsParams
func NewQueryBidsParams(auctionId uint64) QueryBidsParams {
	return QueryBidsParams{
		AuctionId: auctionId,
	}
}

// QueryBidsByBidderParams is the query parameters for 'custom/auction/bids-by-bidder'
type QueryBidsByBidderParams struct {
	Bidder sdk.AccAddress
}

// NewQueryBidsByBidderParams creates a new instance of QueryBidsByBidderParams
func NewQueryBidsByBidderParams(bidder sdk.AccAddress) QueryBidsByBidderParams {
	return QueryBidsByBidderParams{
		Bidder: bidder,
	}
}
