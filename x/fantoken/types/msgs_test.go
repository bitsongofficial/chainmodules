package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	emptyAddr string

	addr1 = sdk.AccAddress(tmhash.SumTruncated([]byte("addr1"))).String()
	addr2 = sdk.AccAddress(tmhash.SumTruncated([]byte("addr2"))).String()
)

// test ValidateBasic for MsgIssueToken
func TestMsgIssueAsset(t *testing.T) {
	addr := sdk.AccAddress(tmhash.SumTruncated([]byte("test"))).String()

	tests := []struct {
		testCase string
		*MsgIssueFanToken
		expectPass bool
	}{
		{"basic good", NewMsgIssueFanToken("stake", "Bitcoin Network", sdk.NewInt(1), "test", addr, sdk.Coin{}), true},
		{"denom empty", NewMsgIssueFanToken("", "Bitcoin Network", sdk.NewInt(1), "test", addr, sdk.Coin{}), false},
		{"denom error", NewMsgIssueFanToken("b&stake", "Bitcoin Network", sdk.NewInt(1), "test", addr, sdk.Coin{}), false},
		{"denom first letter is num", NewMsgIssueFanToken("4stake", "Bitcoin Network", sdk.NewInt(1), "test", addr, sdk.Coin{}), false},
		{"denom too long", NewMsgIssueFanToken("stake123456789012345678901234567890123456789012345678901234567890", "Bitcoin Network", sdk.NewInt(1), "test", addr, sdk.Coin{}), false},
		{"denom too short", NewMsgIssueFanToken("ht", "Bitcoin Network", sdk.NewInt(1), "test", addr, sdk.Coin{}), false},
		{"name empty", NewMsgIssueFanToken("stake", "", sdk.NewInt(1), "test", addr, sdk.Coin{}), false},
		{"name too long", NewMsgIssueFanToken("stake", "Bitcoin Network aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", sdk.NewInt(1), "test", addr, sdk.Coin{}), false},
		{"max supply is zero", NewMsgIssueFanToken("stake", "Bitcoin Network", sdk.ZeroInt(), "test", addr, sdk.Coin{}), true},
	}

	for _, tc := range tests {
		if tc.expectPass {
			require.Nil(t, tc.MsgIssueFanToken.ValidateBasic(), "test: %v", tc.testCase)
		} else {
			require.NotNil(t, tc.MsgIssueFanToken.ValidateBasic(), "test: %v", tc.testCase)
		}
	}
}

// test ValidateBasic for MsgIssueToken
func TestMsgEditFanToken(t *testing.T) {
	owner := sdk.AccAddress(tmhash.SumTruncated([]byte("owner"))).String()
	mintable := false

	tests := []struct {
		testCase string
		*MsgEditFanToken
		expectPass bool
	}{
		{"native basic good", NewMsgEditFanToken("ftbtc", mintable, owner), true},
		{"wrong denom", NewMsgEditFanToken("BT", mintable, owner), false},
		{"loss owner", NewMsgEditFanToken("btc", mintable, ""), false},
	}

	for _, tc := range tests {
		if tc.expectPass {
			require.Nil(t, tc.MsgEditFanToken.ValidateBasic(), "test: %v", tc.testCase)
		} else {
			require.NotNil(t, tc.MsgEditFanToken.ValidateBasic(), "test: %v", tc.testCase)
		}
	}
}

func TestMsgEditFanTokenRoute(t *testing.T) {
	denom := "ubtc"
	mintable := false

	// build a MsgEditToken
	msg := MsgEditFanToken{
		Denom:    denom,
		Mintable: mintable,
	}

	require.Equal(t, "fantoken", msg.Route())
}

func TestMsgEditFanTokenGetSignBytes(t *testing.T) {
	mintable := true

	var msg = MsgEditFanToken{
		Owner:    sdk.AccAddress(tmhash.SumTruncated([]byte("owner"))).String(),
		Denom:    "ubtc",
		Mintable: mintable,
	}

	res := msg.GetSignBytes()

	expected := `{"type":"go-bitsong/token/MsgEditFanToken","value":{"denom":"ubtc","mintable":true,"owner":"cosmos1fsgzj6t7udv8zhf6zj32mkqhcjcpv52ygswxa5"}}`
	require.Equal(t, expected, string(res))
}

func TestMsgMintFanTokenValidateBasic(t *testing.T) {
	testData := []struct {
		msg        string
		denom      string
		owner      string
		recipient  string
		amount     sdk.Int
		expectPass bool
	}{
		{"empty denom", "", addr1, addr2, sdk.NewInt(1000), false},
		{"wrong denom", "bt", addr1, addr2, sdk.NewInt(1000), false},
		{"empty owner", "btc", emptyAddr, addr2, sdk.NewInt(1000), false},
		{"empty to", "ftbtc", addr1, emptyAddr, sdk.NewInt(1000), true},
		{"not empty to", "ftbtc", addr1, addr2, sdk.NewInt(1000), true},
		{"invalid amount", "btc", addr1, addr2, sdk.ZeroInt(), false},
		{"basic good", "ftbtc", addr1, addr2, sdk.NewInt(1000), true},
	}

	for _, td := range testData {
		msg := NewMsgMintFanToken(td.recipient, td.denom, td.owner, td.amount)
		if td.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", td.msg)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", td.msg)
		}
	}
}

func TestMsgBurnFanTokenValidateBasic(t *testing.T) {
	testData := []struct {
		msg        string
		denom      string
		sender     string
		amount     sdk.Int
		expectPass bool
	}{
		{"basic good", "ftbtc", addr1, sdk.NewInt(1000), true},
		{"empty denom", "", addr1, sdk.NewInt(1000), false},
		{"wrong demp,", "bt", addr1, sdk.NewInt(1000), false},
		{"empty sender", "btc", emptyAddr, sdk.NewInt(1000), false},
		{"invalid amount", "btc", addr1, sdk.ZeroInt(), false},
	}

	for _, td := range testData {
		msg := NewMsgBurnFanToken(td.denom, td.sender, td.amount)
		if td.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", td.msg)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", td.msg)
		}
	}
}

func TestMsgTransferFanTokenOwnerValidation(t *testing.T) {
	testData := []struct {
		name       string
		srcOwner   string
		denom      string
		dstOwner   string
		expectPass bool
	}{
		{"empty srcOwner", emptyAddr, "btc", addr1, false},
		{"empty denom", addr1, "", addr2, false},
		{"empty dstOwner", addr1, "btc", emptyAddr, false},
		{"invalid denom", addr1, "btc_min", addr2, false},
		{"basic good", addr1, "ftbtc", addr2, true},
	}

	for _, td := range testData {
		msg := NewMsgTransferFanTokenOwner(td.denom, td.srcOwner, td.dstOwner)
		if td.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", td.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", td.name)
		}
	}
}
