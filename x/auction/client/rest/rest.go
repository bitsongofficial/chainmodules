package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// Rest variable names
// nolint
const (
	RestParamAuctionId = "auction-id"
	RestParamOwner     = "owner"
)

// RegisterHandlers registers token-related REST handlers to a router
func RegisterHandlers(cliCtx client.Context, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

type openAuctionReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	AuctionType int32        `json:"auction_type"`
	NftId       string       `json:"nft_id"`
	Duration    uint64       `json:"duration"`
	MinAmount   string       `json:"min_amount"`
	Owner       string       `json:"owner"` // owner of the token
	limit       uint32       `json:"limit"`
}

type editAuctionReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Duration uint64       `json:"duration"`
	Owner    string       `json:"owner"` //  owner of the token
}
