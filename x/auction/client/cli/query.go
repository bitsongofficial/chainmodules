package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

// GetQueryCmd returns the query commands for the auction module.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                types.ModuleName,
		Short:              "Querying commands for the auction module",
		DisableFlagParsing: true,
	}

	queryCmd.AddCommand(
		GetCmdQueryAuction(),
		GetCmdQueryAllAuctions(),
		GetCmdQueryAuctionsByOwner(),
		GetCmdQueryBid(),
		GetCmdQueryBidsByAuction(),
		GetCmdQueryBidsByBidder(),
	)

	return queryCmd
}

// GetCmdQueryAuction implements the query auction command.
func GetCmdQueryAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auction [id]",
		Long:    "Query an auction by id.",
		Example: fmt.Sprintf("$ %s query auction auction <id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Auction(context.Background(), &types.QueryAuctionRequest{
				Id: auctionId,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Auction)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAllAuctions implements the query all auctions command.
func GetCmdQueryAllAuctions() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auctions",
		Long:    "Query all auctions.",
		Example: fmt.Sprintf("$ %s query auction auctions", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.AllAuctions(
				context.Background(),
				&types.QueryAllAuctionsRequest{
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(res.Auctions)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all auctions")

	return cmd
}

// GetCmdQueryAuctionsByOwner implements the query auctions by owner command.
func GetCmdQueryAuctionsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auctions [owner]",
		Long:    "Query auctions by the owner.",
		Example: fmt.Sprintf("$ %s query auction auctions <owner>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.AuctionsByOwner(
				context.Background(),
				&types.QueryAuctionsByOwnerRequest{
					Owner:      owner.String(),
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(res.Auctions)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all auctions")

	return cmd
}

// GetCmdQueryBid implements the query bid command.
func GetCmdQueryBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bid [auction-id] [bidder]",
		Long:    "Query a bid by auction id and address.",
		Example: fmt.Sprintf("$ %s query auction bid <auction-id> <bidder>", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			bidder, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Bid(context.Background(), &types.QueryBidRequest{
				AuctionId: auctionId,
				Bidder:    bidder.String(),
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Bid)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryBids implements the query bids command.
func GetCmdQueryBidsByAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bids [auction-id]",
		Long:    "Query bids by the auction id.",
		Example: fmt.Sprintf("$ %s query auction bids <auction-id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.BidsByAuction(
				context.Background(),
				&types.QueryBidsByAuctionRequest{
					AuctionId:  auctionId,
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(res.Bids)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all bids")

	return cmd
}

// GetCmdQueryBidsByBidder implements the query bids by bidder command.
func GetCmdQueryBidsByBidder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bids [bidder]",
		Long:    "Query bids by the bidder.",
		Example: fmt.Sprintf("$ %s query auction bids <bidder>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			bidder, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.BidsByBidder(
				context.Background(),
				&types.QueryBidsByBidderRequest{
					Bidder:     bidder.String(),
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(res.Bids)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all bids")

	return cmd
}
