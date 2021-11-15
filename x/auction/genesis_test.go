package auction_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	simapp "github.com/bitsongofficial/chainmodules/app"
	auction "github.com/bitsongofficial/chainmodules/x/auction"
	"github.com/bitsongofficial/chainmodules/x/auction/types"
	nfttypes "github.com/bitsongofficial/chainmodules/x/nft/types"
)

func TestExportGenesis(t *testing.T) {
	app := simapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// export genesis
	genesisState := auction.ExportGenesis(ctx, app.AuctionKeeper)

	require.Equal(t, uint64(0), genesisState.LastAuctionId)
	require.Equal(t, 0, len(genesisState.Auctions))
	require.Equal(t, 0, len(genesisState.Bids))
}

func TestInitGenesis(t *testing.T) {
	app := simapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addr := sdk.AccAddress(tmhash.SumTruncated([]byte("addr1")))

	err := app.NFTKeeper.SetDenom(ctx, nfttypes.NewDenom("testDenom", "testDenom", []string{addr.String()}, []sdk.Dec{sdk.NewDec(100)}, sdk.NewDec(10), addr))
	require.NoError(t, err)
	err = app.NFTKeeper.MintNFT(ctx, "testDenom", "testToken", "testToken", "testTokenURI", addr, addr, true)
	require.NoError(t, err)

	// add auction
	ac := types.NewAuction(0, 0, "testDenom", "testToken", uint64(ctx.BlockTime().Unix()), 86400, sdk.NewCoin("btsg", sdk.NewInt(1000)), addr, 1)

	genesis := types.GenesisState{
		LastAuctionId: 0,
		Auctions:      []types.Auction{ac},
		Bids:          []types.Bid{},
	}

	// initialize genesis
	auction.InitGenesis(ctx, app.AuctionKeeper, genesis)

	// query all auctions
	var acs = app.AuctionKeeper.GetAuctions(ctx, nil)
	require.Equal(t, len(acs), 1)
	require.Equal(t, acs[0], ac)
}
