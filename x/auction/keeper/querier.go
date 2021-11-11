package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryAuction:
			return queryAuction(ctx, req, k, legacyQuerierCdc)
		case types.QueryAllAuctions:
			return queryAllAuctions(ctx, req, k, legacyQuerierCdc)
		case types.QueryAuctionsByOwner:
			return queryAuctionsByOwner(ctx, req, k, legacyQuerierCdc)
		case types.QueryBid:
			return queryBid(ctx, req, k, legacyQuerierCdc)
		case types.QueryBids:
			return queryBids(ctx, req, k, legacyQuerierCdc)
		case types.QueryBidsByBidder:
			return queryBidsByBidder(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown token query endpoint")
		}
	}
}

func queryAuction(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAuctionParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, err
	}

	auction, err := keeper.GetAuctionById(ctx, params.Id)
	if err != nil {
		return nil, err
	}

	return codec.MarshalJSONIndent(legacyQuerierCdc, auction)
}

func queryAllAuctions(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	auctions := keeper.GetAuctions(ctx, nil)
	return codec.MarshalJSONIndent(legacyQuerierCdc, auctions)
}

func queryAuctionsByOwner(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAuctionsByOwnerParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, err
	}
	auctions := keeper.GetAuctions(ctx, params.Owner)
	return codec.MarshalJSONIndent(legacyQuerierCdc, auctions)
}

func queryBid(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryBidParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, err
	}
	bid, err := keeper.GetBid(ctx, params.AuctionId, params.Bidder)
	if err != nil {
		return nil, err
	}
	return codec.MarshalJSONIndent(legacyQuerierCdc, bid)
}

func queryBids(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryBidsParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, err
	}
	bids := keeper.GetBidsByAuctionId(ctx, params.AuctionId)
	return codec.MarshalJSONIndent(legacyQuerierCdc, bids)
}

func queryBidsByBidder(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryBidsByBidderParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, err
	}
	bids := keeper.getBidsByBidder(ctx, params.Bidder)
	return codec.MarshalJSONIndent(legacyQuerierCdc, bids)
}
