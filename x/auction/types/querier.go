package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QueryAuction         = "auction"
	QueryAuctions        = "auctions"
	QueryAuctionsByOwner = "auctions-by-owner"
	QueryBids            = "bids"
)

// QueryAuctionParams is the query parameters for 'custom/auction/auction'
type QueryAuctionParams struct {
	Id uint64
}

// NewQueryAuctionParams creates a new instance of QueryAuctionParams
func NewQueryAuctionParams(id uint64) QueryAuctionParams {
	return QueryAuctionParams{
		Id: id,
	}
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
type QueryBidsParams struct {
	AuctionId uint64
}

// NewQueryBidsParams creates a new instance of QueryBidsParams
func NewQueryBidsParams(auctionId uint64) QueryBidsParams {
	return QueryBidsParams{
		AuctionId: auctionId,
	}
}
