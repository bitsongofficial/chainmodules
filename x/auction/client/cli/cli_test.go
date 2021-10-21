package cli_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tidwall/gjson"

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
	println(123123, txResp.RawLog)
	println(234234, txResp.Logs.String())
	auctionId := gjson.Get(txResp.RawLog, "0.events.1.attributes.0.value").String()

	//------test GetCmdQueryAuctionsByOwner()-------------
	auctions := &[]auctiontypes.Auction{}
	bz, err = auctiontestutil.QueryAuctionsByOwnerExec(clientCtx, from.String())
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.LegacyAmino.UnmarshalJSON(bz.Bytes(), auctions))
	s.Require().Equal(1, len(*auctions))

	//------test GetCmdQueryAllAuctions()-------------
	auctions = &[]auctiontypes.Auction{}
	bz, err = auctiontestutil.QueryAllAuctionsExec(clientCtx, from.String())
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.LegacyAmino.UnmarshalJSON(bz.Bytes(), auctions))
	s.Require().Equal(1, len(*auctions))

	//------test GetCmdQueryAuction()-------------
	var auction *auctiontypes.Auction
	respType = proto.Message(&auctiontypes.Auction{})
	bz, err = auctiontestutil.QueryAuctionExec(clientCtx, auctionId)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType))
	auction = respType.(*auctiontypes.Auction)
	s.Require().Equal("testdenom", auction.NftDenomId)
	s.Require().Equal("testtoken", auction.NftTokenId)

	// ------test GetCmdEditAuction()-------------
	newDuration := uint64(100000)

	args = []string{
		fmt.Sprintf("--%s=%d", auctioncli.FlagDuration, newDuration),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})
	bz, err = auctiontestutil.EditAuctionExec(clientCtx, from.String(), auctionId, args...)

	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&auctiontypes.Auction{})
	bz, err = auctiontestutil.QueryAuctionExec(clientCtx, auctionId)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType))
	auction = respType.(*auctiontypes.Auction)
	s.Require().Equal(newDuration, auction.GetDuration())
}
