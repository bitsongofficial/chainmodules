// Copyright (c) 2016-2021 Shanghai Bianjie AI Technology Inc. (licensed under the Apache License, Version 2.0)
// Modifications Copyright (c) 2021, CRO Protocol Labs ("Crypto.org") (licensed under the Apache License, Version 2.0)
package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTokenName = "name"
	FlagTokenURI  = "uri"
	FlagRecipient = "recipient"
	FlagOwner     = "owner"

	FlagDenomName    = "name"
	FlagDenomID      = "denom-id"
	FlagCreators     = "creators"
	FlagSplitShares  = "split-shares"
	FlagRoyaltyShare = "royalty-share"
)

var (
	FsIssueDenom  = flag.NewFlagSet("", flag.ContinueOnError)
	FsMintNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferNFT = flag.NewFlagSet("", flag.ContinueOnError)
	FsQuerySupply = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryOwner  = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsIssueDenom.String(FlagCreators, "", "Creators of the denom")
	FsIssueDenom.String(FlagSplitShares, "", "Split shares of the denom")
	FsIssueDenom.String(FlagRoyaltyShare, "", "Royalty share of the denom")
	FsIssueDenom.String(FlagDenomName, "", "The name of the denom")

	FsMintNFT.String(FlagTokenURI, "", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsMintNFT.String(FlagRecipient, "", "Receiver of the nft, if not filled, the default is the sender of the transaction")
	FsMintNFT.String(FlagTokenName, "", "The name of the nft")

	FsEditNFT.String(FlagTokenURI, "[do-not-modify]", "URI for the supplemental off-chain token data (should return a JSON object)")
	FsEditNFT.String(FlagTokenName, "[do-not-modify]", "The name of the nft")

	FsQuerySupply.String(FlagOwner, "", "The owner of the nft")

	FsQueryOwner.String(FlagDenomID, "", "The name of the collection")
}
