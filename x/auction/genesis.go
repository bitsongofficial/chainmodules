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

	k.setLastAuctionId(ctx, data.LastAuctionId)
	// init auctions
	for _, auction := range data.Auctions {
		if err := k.addAuction(ctx, auction); err != nil {
			panic(err.Error())
		}
	}

	// init bids
	for _, bid := range data.Bids {
		if err := k.addBid(ctx, bid); err != nil {
			panic(err.Error())
		}
	}
}

// ExportGenesis outputs the genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var auctions []types.Auction
	for _, auction := range k.getAuctions(ctx, nil) {
		t := auction.(*types.Auction)
		auctions = append(auctions, *t)
	}

	var bids []types.Bid
	for _, bid := range k.getBids(ctx, nil) {
		t := bid.(*types.Bid)
		bids = append(bids, *t)
	}

	return &types.GenesisState{
		LastAuctionId: k.getLastAuctionId(),
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
