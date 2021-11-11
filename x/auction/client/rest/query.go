package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router) {
	// Query the auction by id
	r.HandleFunc(fmt.Sprintf("/%s/auction/{%s}", types.ModuleName, RestParamAuctionId), queryAuctionHandlerFn(cliCtx)).Methods("GET")
	// Query all auctions
	r.HandleFunc(fmt.Sprintf("/%s/auctions", types.ModuleName), queryAllAuctionsHandlerFn(cliCtx)).Methods("GET")
	// Query auctions by owner
	r.HandleFunc(fmt.Sprintf("/%s/auctions/owner/{%s}", types.ModuleName, RestParamOwner), queryAuctionsByOwnerHandlerFn(cliCtx)).Methods("GET")
	// Query bid
	r.HandleFunc(fmt.Sprintf("/%s/auction/{%s}/bid/{%s}", types.ModuleName, RestParamAuctionId, RestParamBidder), queryBidHandlerFn(cliCtx)).Methods("GET")
	// Query bids
	r.HandleFunc(fmt.Sprintf("/%s/auction/{%s}/bids", types.ModuleName, RestParamAuctionId), queryBidsHandlerFn(cliCtx)).Methods("GET")
	// Query bids by bidder
	r.HandleFunc(fmt.Sprintf("/%s/bids-by-bidder/{%s}", types.ModuleName, RestParamOwner), queryBidsByBidderHandlerFn(cliCtx)).Methods("GET")
}

// queryAuctionHandlerFn is the HTTP request handler to query auction
func queryAuctionHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		auctionId, err := strconv.ParseUint(vars[RestParamAuctionId], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to parse auction id: %s", vars[RestParamAuctionId]))
			return
		}
		params := types.QueryAuctionParams{
			Id: auctionId,
		}

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAuction), bz,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// queryAllAuctionsHandlerFn is the HTTP request handler to query all auctions
func queryAllAuctionsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := types.QueryAllAuctionsParams{}

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllAuctions), bz,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// queryAuctionsByOwnerHandlerFn is the HTTP request handler to query auctions by owner
func queryAuctionsByOwnerHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ownerStr := r.FormValue(RestParamOwner)

		var err error
		var owner sdk.AccAddress

		if len(ownerStr) > 0 {
			owner, err = sdk.AccAddressFromBech32(ownerStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		params := types.QueryAuctionsByOwnerParams{
			Owner: owner,
		}

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAuctionsByOwner), bz,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// queryBidHandlerFn is the HTTP request handler to query bid by auction id and bidder
func queryBidHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		auctionId, err := strconv.ParseUint(vars[RestParamAuctionId], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to parse auction id: %s", vars[RestParamAuctionId]))
			return
		}
		bidder, err := sdk.AccAddressFromBech32(vars[RestParamBidder])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params := types.QueryBidParams{
			AuctionId: auctionId,
			Bidder:    bidder,
		}

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryBid), bz,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// queryBidsHandlerFn is the HTTP request handler to query bids by auction id
func queryBidsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		auctionId, err := strconv.ParseUint(vars[RestParamAuctionId], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to parse auction id: %s", vars[RestParamAuctionId]))
			return
		}

		params := types.QueryBidsParams{
			AuctionId: auctionId,
		}

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryBids), bz,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// queryBidsByBidderHandlerFn is the HTTP request handler to query bids by bidder
func queryBidsByBidderHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bidderStr := r.FormValue(RestParamBidder)

		var err error
		var bidder sdk.AccAddress

		if len(bidderStr) > 0 {
			bidder, err = sdk.AccAddressFromBech32(bidderStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		params := types.QueryBidsByBidderParams{
			Bidder: bidder,
		}

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryBidsByBidder), bz,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
