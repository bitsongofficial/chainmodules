package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitsongofficial/chainmodules/x/auction/keeper"
	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

// InitGenesis stores the genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	k.SetLastAuctionId(ctx, data.LastAuctionId)
	// init auctions
	for _, auction := range data.Auctions {
		if err := k.AddAuction(ctx, auction); err != nil {
			panic(err.Error())
		}
	}

	// init bids
	for _, bid := range data.Bids {
		if err := k.AddBid(ctx, bid); err != nil {
			panic(err.Error())
		}
	}
}

// ExportGenesis outputs the genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var auctions []types.Auction
	auctions = append(auctions, k.GetAuctions(ctx, nil)...)

	var bids []types.Bid
	bids = append(bids, k.GetAllBids(ctx)...)

	return &types.GenesisState{
		LastAuctionId: k.GetLastAuctionId(ctx),
		Auctions:      auctions,
		Bids:          bids,
	}
}

// DefaultGenesisState returns the default genesis state for testing
func DefaultGenesisState() *types.GenesisState {
	return &types.GenesisState{
		LastAuctionId: 0,
		Auctions:      []types.Auction{},
		Bids:          []types.Bid{},
	}
}
