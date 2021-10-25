package testutil

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	auctioncli "github.com/bitsongofficial/chainmodules/x/auction/client/cli"
)

func OpenAuctionExec(clientCtx client.Context, from string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdOpenAuction(), args)
}

func EditAuctionExec(clientCtx client.Context, from string, auctionId string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		auctionId,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdEditAuction(), args)
}

func CancelAuctionExec(clientCtx client.Context, from string, auctionId string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		auctionId,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdCancelAuction(), args)
}

func OpenBidExec(clientCtx client.Context, from string, auctionId string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		auctionId,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdOpenBid(), args)
}

func CancelBidExec(clientCtx client.Context, from string, auctionId string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		auctionId,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdCancelBid(), args)
}

func WithdrawExec(clientCtx client.Context, from string, auctionId string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		auctionId,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdWithdraw(), args)
}

func QueryAuctionExec(clientCtx client.Context, auctionId string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		auctionId,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdQueryAuction(), args)
}

func QueryAuctionsByOwnerExec(clientCtx client.Context, owner string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		owner,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdQueryAuctionsByOwner(), args)
}

func QueryAllAuctionsExec(clientCtx client.Context, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdQueryAllAuctions(), args)
}

func QueryBidExec(clientCtx client.Context, auctionId string, bidder string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		auctionId,
		bidder,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdQueryBid(), args)
}

func QueryBidsByAuctionExec(clientCtx client.Context, auctionId string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		auctionId,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdQueryBidsByAuction(), args)
}

func QueryBidsByBidderExec(clientCtx client.Context, bidder string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		bidder,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, auctioncli.GetCmdQueryBidsByBidder(), args)
}
