package types

// NewGenesisState creates a new genesis state.
func NewGenesisState(lastAuctionId uint64, auctions []Auction, bids []Bid) GenesisState {
	return GenesisState{
		LastAuctionId: lastAuctionId,
		Auctions:      auctions,
		Bids:          bids,
	}
}

// ValidateGenesis validates the provided auction genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	// validate auction
	for _, auction := range data.Auctions {
		if err := ValidateAuction(auction); err != nil {
			return err
		}
	}

	// validate bid
	for _, bid := range data.Bids {
		if err := ValidateBid(bid); err != nil {
			return err
		}
	}
	return nil
}
