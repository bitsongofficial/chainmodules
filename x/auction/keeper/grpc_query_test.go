package keeper_test

import (
	gocontext "context"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitsongofficial/chainmodules/types"
	auctiontypes "github.com/bitsongofficial/chainmodules/x/auction/types"
)

func (suite *KeeperTestSuite) TestGRPCQueryAuction() {
	app, ctx := suite.app, suite.ctx
	_, _, owner := testdata.KeyTestPubAddr()
	_, _, owner1 := testdata.KeyTestPubAddr()

	auction := auctiontypes.NewAuction(0, 0, "testDenom", "testToken", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	auction1 := auctiontypes.NewAuction(1, 0, "testDenom1", "testToken1", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	auction2 := auctiontypes.NewAuction(2, 0, "testDenom2", "testToken2", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner1, 1)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	auctiontypes.RegisterQueryServer(queryHelper, app.AuctionKeeper)
	queryClient := auctiontypes.NewQueryClient(queryHelper)

	_ = suite.app.AuctionKeeper.AddAuction(ctx, auction)
	_ = suite.app.AuctionKeeper.AddAuction(ctx, auction1)
	_ = suite.app.AuctionKeeper.AddAuction(ctx, auction2)

	// Query Auction
	auctionResp, err := queryClient.Auction(gocontext.Background(), &auctiontypes.QueryAuctionRequest{Id: 0})
	suite.Require().NoError(err)
	suite.Require().NotNil(auctionResp)

	suite.Require().Equal(auctionResp.Auction.NftDenomId, "testDenom")
	suite.Require().Equal(auctionResp.Auction.NftTokenId, "testToken")

	// Query All Auctions
	auctionsResp1, err := queryClient.AllAuctions(gocontext.Background(), &auctiontypes.QueryAllAuctionsRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(auctionsResp1)
	suite.Len(auctionsResp1.Auctions, 3)

	// Query Auctions By Owner
	auctionsResp2, err := queryClient.AuctionsByOwner(gocontext.Background(), &auctiontypes.QueryAuctionsByOwnerRequest{Owner: owner.String()})
	suite.Require().NoError(err)
	suite.Require().NotNil(auctionsResp2)
	suite.Len(auctionsResp2.Auctions, 2)
}

func (suite *KeeperTestSuite) TestGRPCQueryBid() {
	app, ctx := suite.app, suite.ctx
	_, _, owner := testdata.KeyTestPubAddr()
	_, _, bidder := testdata.KeyTestPubAddr()
	_, _, bidder1 := testdata.KeyTestPubAddr()

	auction := auctiontypes.NewAuction(0, 0, "testDenom", "testToken", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 1)
	auction1 := auctiontypes.NewAuction(1, 1, "testDenom1", "testToken1", uint64(suite.ctx.BlockTime().Unix()), 86400, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)), owner, 2)
	bid := auctiontypes.NewBid(0, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)))
	bid1 := auctiontypes.NewBid(1, bidder, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)))
	bid2 := auctiontypes.NewBid(1, bidder1, sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)))

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	auctiontypes.RegisterQueryServer(queryHelper, app.AuctionKeeper)
	queryClient := auctiontypes.NewQueryClient(queryHelper)

	_ = suite.app.AuctionKeeper.AddAuction(ctx, auction)
	_ = suite.app.AuctionKeeper.AddAuction(ctx, auction1)
	_ = suite.app.AuctionKeeper.AddBid(ctx, bid)
	_ = suite.app.AuctionKeeper.AddBid(ctx, bid1)
	_ = suite.app.AuctionKeeper.AddBid(ctx, bid2)

	// Query Bid
	bidResp, err := queryClient.Bid(gocontext.Background(), &auctiontypes.QueryBidRequest{AuctionId: 0, Bidder: bidder.String()})
	suite.Require().NoError(err)
	suite.Require().NotNil(bidResp)

	// Query All Bids
	bidsResp1, err := queryClient.BidsByAuction(gocontext.Background(), &auctiontypes.QueryBidsByAuctionRequest{AuctionId: 0})
	suite.Require().NoError(err)
	suite.Require().NotNil(bidsResp1)
	suite.Len(bidsResp1.Bids, 1)

	bidsResp2, err := queryClient.BidsByAuction(gocontext.Background(), &auctiontypes.QueryBidsByAuctionRequest{AuctionId: 1})
	suite.Require().NoError(err)
	suite.Require().NotNil(bidsResp2)
	suite.Len(bidsResp2.Bids, 2)

	// Query Bids By Owner
	bidsResp3, err := queryClient.BidsByBidder(gocontext.Background(), &auctiontypes.QueryBidsByBidderRequest{Bidder: bidder.String()})
	suite.Require().NoError(err)
	suite.Require().NotNil(bidsResp3)
	suite.Len(bidsResp3.Bids, 2)

	bidsResp4, err := queryClient.BidsByBidder(gocontext.Background(), &auctiontypes.QueryBidsByBidderRequest{Bidder: bidder1.String()})
	suite.Require().NoError(err)
	suite.Require().NotNil(bidsResp4)
	suite.Len(bidsResp4.Bids, 1)
}
