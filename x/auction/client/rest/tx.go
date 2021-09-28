package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func registerTxRoutes(cliCtx client.Context, r *mux.Router) {
	// open an auction
	r.HandleFunc(fmt.Sprintf("/%s/auction", types.ModuleName), openAuctionHandlerFn(cliCtx)).Methods("POST")
	// edit an auction
	r.HandleFunc(fmt.Sprintf("/%s/auction/{%s}", types.ModuleName, RestParamAuctionId), editAuctionHandlerFn(cliCtx)).Methods("PUT")
}

func openAuctionHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req openAuctionReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		minAmount, err := sdk.ParseCoinNormalized(req.MinAmount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to parse min amount: %s", req.MinAmount))
			return
		}

		msg := types.NewMsgOpenAuction(types.AuctionType(req.AuctionType), req.NftId, req.Duration, minAmount, req.Owner, req.limit)

		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}

func editAuctionHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		auctionId, err := strconv.ParseUint(vars[RestParamAuctionId], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to parse auction id: %s", vars[RestParamAuctionId]))
			return
		}

		var req editAuctionReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgEditAuction(auctionId, req.Duration, req.Owner)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}
