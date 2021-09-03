// Copyright (c) 2016-2021 Shanghai Bianjie AI Technology Inc. (licensed under the Apache License, Version 2.0)
// Modifications Copyright (c) 2021, CRO Protocol Labs ("Crypto.org") (licensed under the Apache License, Version 2.0)
package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// RegisterHandlers registers the NFT REST routes.
func RegisterHandlers(cliCtx client.Context, r *mux.Router, queryRoute string) {
	registerQueryRoutes(cliCtx, r, queryRoute)
	registerTxRoutes(cliCtx, r, queryRoute)
}

const (
	RestParamDenomID   = "denom-id"
	RestParamDenomName = "denom-name"
	RestParamTokenID   = "token-id"
	RestParamOwner     = "owner"
)

type issueDenomReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Owner         string       `json:"owner"`
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Creators      []string     `json:"creators"`
	SplitShares   []sdk.Dec    `json:"split_shares"`
	RoyaltyShares []sdk.Dec    `json:"royalty_shares"`
}

type mintNFTReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Owner     string       `json:"owner"`
	Recipient string       `json:"recipient"`
	DenomID   string       `json:"denom_id"`
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	URI       string       `json:"uri"`
	Data      string       `json:"data"`
}

type editNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner   string       `json:"owner"`
	Name    string       `json:"name"`
	Data    string       `json:"data"`
}

type transferNFTReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Owner     string       `json:"owner"`
	Recipient string       `json:"recipient"`
	Name      string       `json:"name"`
	URI       string       `json:"uri"`
	Data      string       `json:"data"`
}

type burnNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner   string       `json:"owner"`
}
