package keeper

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Auction(c context.Context, req *types.QueryAuctionRequest) (*types.QueryAuctionResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	auction, err := k.GetAuctionById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "auction %d not found", req.Id)
	}

	return &types.QueryAuctionResponse{Auction: &auction}, nil
}

func (k Keeper) AllAuctions(c context.Context, req *types.QueryAllAuctionsRequest) (*types.QueryAllAuctionsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var auctions []types.Auction
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	auctionStore := prefix.NewStore(store, types.PrefixAuctionById)
	pageRes, err := query.Paginate(auctionStore, req.Pagination, func(key []byte, value []byte) error {
		var auction types.Auction
		k.cdc.MustUnmarshalBinaryBare(value, &auction)
		auctions = append(auctions, auction)
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryAllAuctionsResponse{Auctions: auctions, Pagination: pageRes}, nil
}

func (k Keeper) AuctionsByOwner(c context.Context, req *types.QueryAuctionsByOwnerRequest) (*types.QueryAuctionsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var owner sdk.AccAddress
	var err error
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil || owner == nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid owner address (%s)", err))
		}
	}

	var auctions []types.Auction
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	auctionStore := prefix.NewStore(store, append(types.PrefixAuctionsByOwner, owner.Bytes()...))
	pageRes, err = query.Paginate(auctionStore, req.Pagination, func(key []byte, value []byte) error {
		var auctionId gogotypes.UInt64Value
		k.cdc.MustUnmarshalBinaryBare(value, &auctionId)
		auction, err := k.GetAuctionById(ctx, auctionId.Value)
		if err == nil {
			auctions = append(auctions, auction)
		}
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryAuctionsByOwnerResponse{Auctions: auctions, Pagination: pageRes}, nil
}

func (k Keeper) Bid(c context.Context, req *types.QueryBidRequest) (*types.QueryBidResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var bidder sdk.AccAddress
	var err error
	if len(req.Bidder) > 0 {
		bidder, err = sdk.AccAddressFromBech32(req.Bidder)
		if err != nil || bidder == nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid bidder address (%s)", err))
		}
	}

	bid, err := k.GetBid(ctx, req.AuctionId, bidder)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "bid not found")
	}

	return &types.QueryBidResponse{Bid: bid}, nil
}

func (k Keeper) BidsByAuction(c context.Context, req *types.QueryBidsByAuctionRequest) (*types.QueryBidsByAuctionResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var bids []types.Bid
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	bidStore := prefix.NewStore(store, types.KeyBidsByAuctionId(req.AuctionId))
	pageRes, err := query.Paginate(bidStore, req.Pagination, func(key []byte, value []byte) error {
		var bid types.Bid
		k.cdc.MustUnmarshalBinaryBare(value, &bid)
		bids = append(bids, bid)
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryBidsByAuctionResponse{Bids: bids, Pagination: pageRes}, nil
}

func (k Keeper) BidsByBidder(c context.Context, req *types.QueryBidsByBidderRequest) (*types.QueryBidsByBidderResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var bidder sdk.AccAddress
	var err error
	if len(req.Bidder) > 0 {
		bidder, err = sdk.AccAddressFromBech32(req.Bidder)
		if err != nil || bidder == nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid bidder address (%s)", err))
		}
	}

	var bids []types.Bid
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	bidStore := prefix.NewStore(store, append(types.PrefixBidsByBidder, bidder.Bytes()...))
	pageRes, err = query.Paginate(bidStore, req.Pagination, func(key []byte, value []byte) error {
		var auctionId gogotypes.UInt64Value
		k.cdc.MustUnmarshalBinaryBare(value, &auctionId)
		bid, err := k.GetBid(ctx, auctionId.Value, bidder)
		if err == nil {
			bids = append(bids, bid)
		}
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryBidsByBidderResponse{Bids: bids, Pagination: pageRes}, nil
}
