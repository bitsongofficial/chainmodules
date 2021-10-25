package rest_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"

	simapp "github.com/bitsongofficial/chainmodules/app"
	auctioncli "github.com/bitsongofficial/chainmodules/x/auction/client/cli"
	auctiontestutil "github.com/bitsongofficial/chainmodules/x/auction/client/testutil"
	auctiontypes "github.com/bitsongofficial/chainmodules/x/auction/types"
	nftcli "github.com/bitsongofficial/chainmodules/x/nft/client/cli"
	nfttestutil "github.com/bitsongofficial/chainmodules/x/nft/client/testutil"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := simapp.NewConfig()
	cfg.NumValidators = 1

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestAuction() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	baseURL := val.APIAddress
	// -----Issue NFT Denom--------------------------------------------------------
	from := val.Address
	splitShares := "100"
	royaltyShare := "10"

	args := []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagDenomName, "testdenom"),
		fmt.Sprintf("--%s=%s", nftcli.FlagCreators, from.String()),
		fmt.Sprintf("--%s=%s", nftcli.FlagSplitShares, splitShares),
		fmt.Sprintf("--%s=%s", nftcli.FlagRoyaltyShare, royaltyShare),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType := proto.Message(&sdk.TxResponse{})
	expectedCode := uint32(0)

	bz, err := nfttestutil.IssueDenomExec(val.ClientCtx, from.String(), "testdenom", args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp := respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	// -----Mint NFT Token--------------------------------------------------------
	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagRecipient, from.String()),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenURI, "testTokenURI"),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, "testtoken"),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})

	bz, err = nfttestutil.MintNFTExec(val.ClientCtx, from.String(), "testdenom", "testtoken", args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	// ---------------------------------------------------------------------------
	auctionType := 0
	nftDenomId := "testdenom"
	nftTokenId := "testtoken"
	duration := 86400
	minAmount := "1000000ubtsg"
	limit := 1

	//------test GetCmdOpenAuction()-------------
	args = []string{
		fmt.Sprintf("--%s=%d", auctioncli.FlagAuctionType, auctionType),
		fmt.Sprintf("--%s=%s", auctioncli.FlagNftDenomId, nftDenomId),
		fmt.Sprintf("--%s=%s", auctioncli.FlagNftTokenId, nftTokenId),
		fmt.Sprintf("--%s=%d", auctioncli.FlagDuration, duration),
		fmt.Sprintf("--%s=%s", auctioncli.FlagMinAmount, minAmount),
		fmt.Sprintf("--%s=%d", auctioncli.FlagLimit, limit),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	bz, err = auctiontestutil.OpenAuctionExec(clientCtx, from.String(), args...)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)
	auctionId := gjson.Get(txResp.RawLog, "0.events.1.attributes.0.value").String()

	//------test GetCmdQueryAuction()-------------
	url := fmt.Sprintf("%s/bitsong/auction/v1beta1/auction/%s", baseURL, auctionId)
	resp, err := rest.GetRequest(url)
	respType = proto.Message(&auctiontypes.QueryAuctionResponse{})
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(resp, respType))
	auctionResp := respType.(*auctiontypes.QueryAuctionResponse)
	auction := auctionResp.Auction
	s.Require().NoError(err)
	s.Require().Equal("testdenom", auction.NftDenomId)
	s.Require().Equal("testtoken", auction.NftTokenId)

	//------test GetCmdQueryAllAuctions()-------------
	url = fmt.Sprintf("%s/bitsong/auction/v1beta1/auctions", baseURL)
	resp, err = rest.GetRequest(url)
	respType = proto.Message(&auctiontypes.QueryAllAuctionsResponse{})
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(resp, respType))
	auctionsResp := respType.(*auctiontypes.QueryAllAuctionsResponse)
	s.Require().Equal(1, len(auctionsResp.Auctions))

	//------test GetCmdQueryAuctionsByOwner()-------------
	url = fmt.Sprintf("%s/bitsong/auction/v1beta1/auctions/%s", baseURL, from.String())
	resp, err = rest.GetRequest(url)
	respType = proto.Message(&auctiontypes.QueryAuctionsByOwnerResponse{})
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(resp, respType))
	auctionsResp1 := respType.(*auctiontypes.QueryAuctionsByOwnerResponse)
	s.Require().Equal(1, len(auctionsResp1.Auctions))

	//------setup test env for bid test-------------
	info, _, err := clientCtx.Keyring.NewMnemonic("NewBidderAddr", keyring.English, sdk.FullFundraiserPath, hd.Secp256k1)
	s.Require().NoError(err)

	bidder := sdk.AccAddress(info.GetPubKey().Address())

	_, err = banktestutil.MsgSendExec(
		clientCtx,
		from,
		bidder,
		sdk.NewCoins(sdk.NewInt64Coin(s.cfg.BondDenom, 2000000)),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)

	// ------test GetCmdOpenBid()-------------
	bidAmount := "1000000ubtsg"
	args = []string{
		fmt.Sprintf("--%s=%s", auctioncli.FlagBidAmount, bidAmount),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})
	bz, err = auctiontestutil.OpenBidExec(clientCtx, bidder.String(), auctionId, args...)

	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	//------test GetCmdQueryBid()-------------
	url = fmt.Sprintf("%s/bitsong/auction/v1beta1/bid/%s/%s", baseURL, auctionId, bidder.String())
	resp, err = rest.GetRequest(url)
	respType = proto.Message(&auctiontypes.QueryBidResponse{})
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(resp, respType))
	bidResp := respType.(*auctiontypes.QueryBidResponse)
	bid := bidResp.Bid
	s.Require().NoError(err)
	s.Require().Equal(auctionId, strconv.FormatUint(bid.AuctionId, 10))
	s.Require().Equal(bidder.String(), bid.Bidder)

	//------test GetCmdQueryBidsByAuction()-------------
	url = fmt.Sprintf("%s/bitsong/auction/v1beta1/bids/%s", baseURL, auctionId)
	resp, err = rest.GetRequest(url)
	respType = proto.Message(&auctiontypes.QueryBidsByAuctionResponse{})
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(resp, respType))
	bidsResp := respType.(*auctiontypes.QueryBidsByAuctionResponse)
	s.Require().Equal(1, len(bidsResp.Bids))

	//------test GetCmdQueryBidsByBidder()-------------
	url = fmt.Sprintf("%s/bitsong/auction/v1beta1/bids-by-bidder/%s", baseURL, bidder.String())
	resp, err = rest.GetRequest(url)
	respType = proto.Message(&auctiontypes.QueryBidsByBidderResponse{})
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(resp, respType))
	bidsResp1 := respType.(*auctiontypes.QueryBidsByBidderResponse)
	s.Require().Equal(1, len(bidsResp1.Bids))
}
