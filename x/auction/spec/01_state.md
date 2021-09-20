# State

## Auction
Definition of data structure of Auction
Auction can be opened through `MsgOpenAuction`, or you can edit through `MsgEditAuction`. Also it is possible to cancel the auction through `MsgCancelAuction`.
The winner can get the NFT through `MsgWithdraw` after the auction is ended.

```go
type Auction struct {
	Id				uint64		// Auction Id
	AuctionType		uint32		// Auction	Type
	NftId			string		// NFTId which is auctioned
	StartTime		uint64		// Auction start time
	Duration 		uint64		// Auction Duration
	MinAmount		sdk.Coin	// Bid Min Amount
	Owner			string		// Auction Creator
	Limit			uint32		// Auction Winner Limit
}
```

## Bid
Definition of data structure of Bid
Bid can be opened through `MsgOpenBid`, or you can cancel it through `MsgCancelBid`.

```go
type Bid struct {
	AuctionId	uint64		// AuctionId where to bid
	Bidder		string		// Address where to bid
	BidAmount	sdk.Coin	// Bid Amount
}
```