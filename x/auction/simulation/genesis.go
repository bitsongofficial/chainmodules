package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"

	auctiontypes "github.com/bitsongofficial/chainmodules/x/auction/types"
)

// RandomizedGenState generates a random GenesisState for bank
func RandomizedGenState(simState *module.SimulationState) {
	var auctions []auctiontypes.Auction
	var bids []auctiontypes.Bid

	simState.AppParams.GetOrGenerate(
		simState.Cdc, "", nil, simState.Rand,
		func(r *rand.Rand) {
			for i := 0; i < 5; i++ {
				auctions = append(auctions, randAuction(r, simState.Accounts))
				bids = append(bids, randBid(r, auctions[i].Id, simState.Accounts))
			}
		},
	)

	auctionGenesis := auctiontypes.NewGenesisState(
		0, auctions, bids,
	)

	bz, err := json.MarshalIndent(&auctionGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", auctiontypes.ModuleName, bz)

	simState.GenState[auctiontypes.ModuleName] = simState.Cdc.MustMarshalJSON(&auctionGenesis)
}
