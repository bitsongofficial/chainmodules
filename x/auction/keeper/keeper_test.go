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
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

const (
	isCheckTx = false
)

var (
	owner        = sdk.AccAddress(tmhash.SumTruncated([]byte("auction owner")))
	bidder       = sdk.AccAddress(tmhash.SumTruncated([]byte("bidder")))
	bidder1      = sdk.AccAddress(tmhash.SumTruncated([]byte("bidder1")))
	bidder2      = sdk.AccAddress(tmhash.SumTruncated([]byte("bidder2")))
	initAmt      = sdk.NewIntWithDecimal(100000000, int(6))
	initCoin     = sdk.Coins{sdk.NewCoin(types.BondDenom, initAmt)}
	initMintCoin = sdk.Coins{sdk.NewCoin(types.BondDenom, initAmt.Mul(sdk.NewInt(10)))}
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

	// init tokens to addr
	err = suite.bk.MintCoins(suite.ctx, minttypes.ModuleName, initMintCoin)
	suite.NoError(err)
	err = suite.bk.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, initCoin)
	suite.NoError(err)
	err = suite.bk.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder1, initCoin)
	suite.NoError(err)
	err = suite.bk.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder2, initCoin)
	suite.NoError(err)
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestOpenAuction() {
	newAuction := auctiontypes.NewAuction(1, 0, "testDenom", "testToken", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	id, err := suite.keeper.OpenAuction(suite.ctx, 0, "testDenom", "testToken", 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	suite.NoError(err)
	suite.True(suite.keeper.HasAuction(suite.ctx, id))

	nft, err := suite.nk.GetNFT(suite.ctx, "testDenom", "testToken")
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), suite.ak.GetModuleAddress(auctiontypes.ModuleName).String())

	auction, err := suite.keeper.GetAuctionById(suite.ctx, id)
	suite.NoError(err)
	suite.EqualValues(auction, newAuction)
}

func (suite *KeeperTestSuite) TestEditAuction() {
	newAuction := auctiontypes.NewAuction(0, 0, "testDenom", "testToken", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	suite.keeper.SetAuction(suite.ctx, newAuction)

	err := suite.keeper.EditAuction(suite.ctx, newAuction.Id, 100000, owner)
	suite.NoError(err)

	auction, err := suite.keeper.GetAuctionById(suite.ctx, 0)
	suite.NoError(err)
	suite.EqualValues(auction.Duration, 100000)
}

func (suite *KeeperTestSuite) TestCancelAuction() {
	newAuction := auctiontypes.NewAuction(0, 0, "testDenom", "testToken", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	suite.keeper.SetAuction(suite.ctx, newAuction)

	err := suite.keeper.CancelAuction(suite.ctx, newAuction.Id, owner)
	suite.NoError(err)

	auction, err := suite.keeper.GetAuctionById(suite.ctx, 0)
	suite.NoError(err)
	suite.EqualValues(auction.Cancelled, true)
}

func (suite *KeeperTestSuite) TestOpenBid() {
	newAuction := auctiontypes.NewAuction(0, 0, "testDenom", "testToken", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	suite.keeper.SetAuction(suite.ctx, newAuction)

	err := suite.keeper.OpenBid(suite.ctx, 1, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)))
	suite.EqualError(err, "auction 1 does not exist: auction does not exist")

	err = suite.keeper.OpenBid(suite.ctx, 0, bidder, sdk.NewCoin("newCoinDenom", sdk.NewInt(1000)))
	suite.EqualError(err, "bid amount denom is different with auction min amount denom: invalid bid amount denom")

	err = suite.keeper.OpenBid(suite.ctx, 0, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(999)))
	suite.EqualError(err, "the bid amount is less than the auction min amount: bid amount is not enough to be a winner")

	err = suite.keeper.OpenBid(suite.ctx, 0, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(1111)))
	suite.NoError(err)

	bid, err := suite.keeper.GetBid(suite.ctx, 0, bidder)
	suite.NoError(err)
	suite.Equal(bid.BidAmount.Amount, sdk.NewInt(1111))

	err = suite.keeper.OpenBid(suite.ctx, 0, bidder1, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)))
	suite.EqualError(err, "the bid amount is not enough to be a winner: bid amount is not enough to be a winner")

	err = suite.keeper.OpenBid(suite.ctx, 0, bidder1, sdk.NewCoin(types.BondDenom, sdk.NewInt(2000)))
	suite.NoError(err)

	bids := suite.keeper.GetBidsByAuctionId(suite.ctx, 0)
	suite.Equal(len(bids), 1)
	suite.Equal(bids[0].Bidder, bidder1.String())
}

func (suite *KeeperTestSuite) TestCancelBid() {
	newAuction := auctiontypes.NewAuction(0, 0, "testDenom", "testToken", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	suite.keeper.SetAuction(suite.ctx, newAuction)

	err := suite.keeper.CancelBid(suite.ctx, 0, bidder)
	suite.Error(err) // bid not exists

	err = suite.keeper.OpenBid(suite.ctx, 0, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(1111)))
	suite.NoError(err)

	err = suite.keeper.CancelBid(suite.ctx, 0, bidder)
	suite.NoError(err)

	_, err = suite.keeper.GetBid(suite.ctx, 0, bidder)
	suite.Error(err) // bid not exists
}

func (suite *KeeperTestSuite) TestWithdrawSingleEdition() {
	id, err := suite.keeper.OpenAuction(suite.ctx, 0, "testDenom", "testToken", 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	suite.NoError(err)

	err = suite.keeper.OpenBid(suite.ctx, id, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)))
	suite.NoError(err)

	err = suite.keeper.Withdraw(suite.ctx, id, bidder)
	suite.EqualError(err, "the auction is still in progress or cancelled: invalid auction")

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour * 25))

	err = suite.keeper.Withdraw(suite.ctx, id, bidder1)
	suite.EqualError(err, "bid not found: bid does not exist")

	err = suite.keeper.Withdraw(suite.ctx, id, bidder)
	suite.NoError(err)

	nft, err := suite.nk.GetNFT(suite.ctx, "testDenom", "testToken")
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), bidder.String())

	err = suite.keeper.Withdraw(suite.ctx, id, owner)
	suite.NoError(err)
	amount := suite.bk.GetBalance(suite.ctx, owner, types.BondDenom)
	suite.Equal(amount.Amount, sdk.NewInt(1000))

	nft, err = suite.nk.GetNFT(suite.ctx, "testDenom", "testToken")
	suite.NoError(err)
	suite.Equal(nft.GetPrimaryStatus(), false)
}

func (suite *KeeperTestSuite) TestWithdrawOpenEdition() {
	id, err := suite.keeper.OpenAuction(suite.ctx, 1, "testDenom", "testToken", 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 0)
	suite.NoError(err)

	err = suite.keeper.OpenBid(suite.ctx, id, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)))
	suite.NoError(err)

	err = suite.keeper.OpenBid(suite.ctx, id, bidder1, sdk.NewCoin(types.BondDenom, sdk.NewInt(1001)))
	suite.NoError(err)

	err = suite.keeper.OpenBid(suite.ctx, id, bidder2, sdk.NewCoin(types.BondDenom, sdk.NewInt(0)))
	suite.NoError(err)

	err = suite.keeper.Withdraw(suite.ctx, id, bidder)
	suite.EqualError(err, "the auction is still in progress or cancelled: invalid auction")

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour * 25))

	err = suite.keeper.Withdraw(suite.ctx, id, bidder)
	suite.NoError(err)

	err = suite.keeper.Withdraw(suite.ctx, id, bidder1)
	suite.NoError(err)

	err = suite.keeper.Withdraw(suite.ctx, id, bidder2)
	suite.NoError(err)

	err = suite.keeper.Withdraw(suite.ctx, id, owner)
	suite.NoError(err)
	amount := suite.bk.GetBalance(suite.ctx, owner, types.BondDenom)
	suite.Equal(amount.Amount, sdk.NewInt(2001))

	nft, err := suite.nk.GetNFT(suite.ctx, "testDenom", "testToken")
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), owner.String())
	suite.Equal(nft.GetPrimaryStatus(), false)

	nft, err = suite.nk.GetNFT(suite.ctx, "testDenom", "testToken"+bidder.String())
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), bidder.String())
	suite.Equal(nft.GetPrimaryStatus(), false)

	nft, err = suite.nk.GetNFT(suite.ctx, "testDenom", "testToken"+bidder1.String())
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), bidder1.String())
	suite.Equal(nft.GetPrimaryStatus(), false)

	nft, err = suite.nk.GetNFT(suite.ctx, "testDenom", "testToken"+bidder2.String())
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), bidder2.String())
	suite.Equal(nft.GetPrimaryStatus(), false)
}

func (suite *KeeperTestSuite) TestWithdrawLimitedEdition() {
	id, err := suite.keeper.OpenAuction(suite.ctx, 2, "testDenom", "testToken", 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 2)
	suite.NoError(err)

	err = suite.keeper.OpenBid(suite.ctx, id, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)))
	suite.NoError(err)

	err = suite.keeper.OpenBid(suite.ctx, id, bidder1, sdk.NewCoin(types.BondDenom, sdk.NewInt(1001)))
	suite.NoError(err)

	err = suite.keeper.OpenBid(suite.ctx, id, bidder2, sdk.NewCoin(types.BondDenom, sdk.NewInt(1002)))
	suite.NoError(err)

	err = suite.keeper.Withdraw(suite.ctx, id, bidder1)
	suite.EqualError(err, "the auction is still in progress or cancelled: invalid auction")

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour * 25))

	err = suite.keeper.Withdraw(suite.ctx, id, bidder)
	suite.EqualError(err, "bid not found: bid does not exist")

	err = suite.keeper.Withdraw(suite.ctx, id, bidder1)
	suite.NoError(err)

	err = suite.keeper.Withdraw(suite.ctx, id, bidder2)
	suite.NoError(err)

	err = suite.keeper.Withdraw(suite.ctx, id, owner)
	suite.NoError(err)
	amount := suite.bk.GetBalance(suite.ctx, owner, types.BondDenom)
	suite.Equal(amount.Amount, sdk.NewInt(2003))

	nft, err := suite.nk.GetNFT(suite.ctx, "testDenom", "testToken")
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), owner.String())
	suite.Equal(nft.GetPrimaryStatus(), false)

	nft, err = suite.nk.GetNFT(suite.ctx, "testDenom", "testToken"+bidder1.String())
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), bidder1.String())
	suite.Equal(nft.GetPrimaryStatus(), false)

	nft, err = suite.nk.GetNFT(suite.ctx, "testDenom", "testToken"+bidder2.String())
	suite.NoError(err)
	suite.Equal(nft.GetOwner().String(), bidder2.String())
	suite.Equal(nft.GetPrimaryStatus(), false)
}
