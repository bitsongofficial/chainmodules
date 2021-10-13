package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	simapp "github.com/bitsongofficial/chainmodules/app"
	"github.com/bitsongofficial/chainmodules/types"
	"github.com/bitsongofficial/chainmodules/x/auction/keeper"
	auctiontypes "github.com/bitsongofficial/chainmodules/x/auction/types"
	nftkeeper "github.com/bitsongofficial/chainmodules/x/nft/keeper"
	nfttypes "github.com/bitsongofficial/chainmodules/x/nft/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

const (
	isCheckTx = false
)

var (
	owner = sdk.AccAddress(tmhash.SumTruncated([]byte("auctionTest")))
)

type KeeperTestSuite struct {
	suite.Suite

	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	keeper      keeper.Keeper
	ak          authkeeper.AccountKeeper
	bk          bankkeeper.Keeper
	nk          nftkeeper.Keeper
	app         *simapp.Bitsong
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(isCheckTx)

	suite.legacyAmino = app.LegacyAmino()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.keeper = app.AuctionKeeper
	suite.ak = app.AccountKeeper
	suite.bk = app.BankKeeper
	suite.nk = app.NFTKeeper
	suite.app = app

	// init nfts to addr
	err := suite.nk.SetDenom(suite.ctx, nfttypes.NewDenom("testDenom", "testDenom", []string{owner.String()}, []sdk.Dec{sdk.NewDec(100)}, sdk.NewDec(10), owner))
	suite.NoError(err)
	err = suite.nk.MintNFT(suite.ctx, "testDenom", "testToken", "testToken", "testTokenURI", owner, owner, true)
	suite.NoError(err)
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestOpenAuction() {
	newAuction := auctiontypes.NewAuction(0, 0, "testDenom", "testToken", uint64(time.Now().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	_, err := suite.keeper.OpenAuction(suite.ctx, 0, "testDenom", "testToken", 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	suite.NoError(err)
	suite.True(suite.keeper.HasAuction(suite.ctx, 0))

	nft, err := suite.nk.GetNFT(suite.ctx, "testDenom", "testToken")
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), suite.ak.GetModuleAddress(auctiontypes.ModuleName).String())

	auction, err := suite.keeper.GetAuctionById(suite.ctx, 0)
	suite.NoError(err)
	suite.EqualValues(auction, newAuction)
}

func (suite *KeeperTestSuite) TestEditAuction() {
	newAuction := auctiontypes.NewAuction(0, 0, "testDenom", "testToken", uint64(time.Now().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	suite.keeper.SetAuction(suite.ctx, newAuction)

	err := suite.keeper.EditAuction(suite.ctx, newAuction.Id, 100000, owner)
	suite.NoError(err)

	auction, err := suite.keeper.GetAuctionById(suite.ctx, 0)
	suite.NoError(err)
	suite.EqualValues(auction.Duration, 100000)
}
