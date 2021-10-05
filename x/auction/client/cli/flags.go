package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagAuctionType = "auction-type"
	FlagNftDenomId  = "nft-denom-id"
	FlagNftTokenId  = "nft-token-id"
	FlagDuration    = "duration"
	FlagMinAmount   = "min-amount"
	FlagLimit       = "limit"
	FlagBidAmount   = "bid-amount"
)

var (
	FsOpenAuction   = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditAuction   = flag.NewFlagSet("", flag.ContinueOnError)
	FsCancelAuction = flag.NewFlagSet("", flag.ContinueOnError)
	FsOpenBid       = flag.NewFlagSet("", flag.ContinueOnError)
	FsCancelBid     = flag.NewFlagSet("", flag.ContinueOnError)
	FsWithdraw      = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsOpenAuction.Int32(FlagAuctionType, 0, "The auction type 0 ~ 2")
	FsOpenAuction.String(FlagNftDenomId, "", "The nft denom id which is auctioned")
	FsOpenAuction.String(FlagNftTokenId, "", "The nft token id which is auctioned")
	FsOpenAuction.Uint64(FlagDuration, 0, "The auction duration")
	FsOpenAuction.String(FlagMinAmount, "", "The auction min amount")
	FsOpenAuction.Uint32(FlagLimit, 0, "The auction winner limit")

	FsEditAuction.Uint64(FlagDuration, 0, "The auction duration")

	FsOpenBid.String(FlagBidAmount, "", "The bid amount")
}
