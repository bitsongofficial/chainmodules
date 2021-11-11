# State

## Auction
Definition of data structure of Auction
Auction can be opened through `MsgOpenAuction`, or you can edit through `MsgEditAuction`. Also it is possible to cancel the auction through `MsgCancelAuction`.
The winner can get the NFT and the auction owner can withdraw the tokens through `MsgWithdraw` after the auction is ended.

```go
type Auction struct {
	Id				uint64			// Auction Id
	AuctionType		AuctionType		// Auction	Type
	NftDenomId		string			// Denom Id of the NFT which is auctioned
	NftTokenId		string			// Token Id of the NFT which is auctioned
	StartTime		uint64			// Auction start time
	Duration 		uint64			// Auction Duration
	MinAmount		sdk.Coin		// Bid Min Amount
	Owner			string			// Auction Creator
	Limit			uint32			// Auction Winner Limit
	Cancelled		bool			// Cancel status of the Auction
}
```

AuctionType is the type of an auction.

```go
type AuctionType int32

const (
	Single_Edition  AuctionType = 0
	Open_Edition    AuctionType = 1
	Limited_Edition AuctionType = 2
)
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