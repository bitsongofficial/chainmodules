package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

// NewTxCmd returns the transaction commands for the token module.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Auction transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdOpenAuction(),
		GetCmdEditAuction(),
		GetCmdCancelAuction(),
		GetCmdOpenBid(),
		GetCmdCancelBid(),
		GetCmdWithdraw(),
	)

	return txCmd
}

// GetCmdOpenAuction implements the open auction command
func GetCmdOpenAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "open-auction",
		Long: "open a new auction",
		Example: fmt.Sprintf(
			"$ %s tx auction open-auction "+
				"--auction-type=<auction-type> "+
				"--nft-id=<nft-id> "+
				"--duration=\"86400\" "+
				"--min-amount=\"1000000ubtsg\" "+
				"--limit=<limit> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()
			auctionType, err := cmd.Flags().GetInt32(FlagAuctionType)
			if err != nil {
				return err
			}
			nftId, err := cmd.Flags().GetString(FlagNftId)
			if err != nil {
				return err
			}
			duration, err := cmd.Flags().GetUint64(FlagDuration)
			if err != nil {
				return err
			}
			minAmountStr, err := cmd.Flags().GetString(FlagMinAmount)
			if err != nil {
				return err
			}
			minAmount, err := sdk.ParseCoinNormalized(minAmountStr)
			if err != nil {
				return fmt.Errorf("failed to parse min amount: %s", minAmountStr)
			}
			limit, err := cmd.Flags().GetUint32(FlagLimit)
			if err != nil {
				return err
			}

			msg := types.NewMsgOpenAuction(types.AuctionType(auctionType), nftId, duration, minAmount, owner.String(), limit)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsOpenAuction)
	_ = cmd.MarkFlagRequired(FlagAuctionType)
	_ = cmd.MarkFlagRequired(FlagNftId)
	_ = cmd.MarkFlagRequired(FlagDuration)
	_ = cmd.MarkFlagRequired(FlagMinAmount)
	_ = cmd.MarkFlagRequired(FlagLimit)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditAuction implements the edit auction command
func GetCmdEditAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "edit-auction",
		Long: "Edit an existing auction",
		Example: fmt.Sprintf(
			"$ %s tx auction edit-auction [id] "+
				"--duration=\"86400\" "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()

			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			duration, err := cmd.Flags().GetUint64(FlagDuration)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditAuction(auctionId, duration, owner.String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsEditAuction)
	_ = cmd.MarkFlagRequired(FlagDuration)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdCancelAuction implements the cancel auction command
func GetCmdCancelAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "cancel-auction",
		Long: "Cancel an existing auction",
		Example: fmt.Sprintf(
			"$ %s tx auction cancel-auction [id] "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()

			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelAuction(auctionId, owner.String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsCancelAuction)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdOpenBid implements the open bid command
func GetCmdOpenBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "open-bid",
		Long: "open a new bid",
		Example: fmt.Sprintf(
			"$ %s tx auction open-bid [auction-id]"+
				"--bid-amount=\"1000000ubtsg\" "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bidder := clientCtx.GetFromAddress()
			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			bidAmountStr, err := cmd.Flags().GetString(FlagBidAmount)
			if err != nil {
				return err
			}
			bidAmount, err := sdk.ParseCoinNormalized(bidAmountStr)
			if err != nil {
				return fmt.Errorf("failed to parse min amount: %s", bidAmountStr)
			}

			msg := types.NewMsgOpenBid(auctionId, bidder.String(), bidAmount)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsOpenBid)
	_ = cmd.MarkFlagRequired(FlagBidAmount)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdCancelBid implements the cancel bid command
func GetCmdCancelBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "cancel-bid",
		Long: "cancel an existing bid",
		Example: fmt.Sprintf(
			"$ %s tx auction cancel-bid [auction-id]"+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bidder := clientCtx.GetFromAddress()

			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelBid(auctionId, bidder.String())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsOpenBid)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdWithdraw implements the withdraw command
func GetCmdWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw",
		Long: "withdraw",
		Example: fmt.Sprintf(
			"$ %s tx auction withdraw [auction-id]"+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient := clientCtx.GetFromAddress()
			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(auctionId, recipient.String())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsOpenBid)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
