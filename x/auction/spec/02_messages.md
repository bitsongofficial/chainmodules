# Messages

## MsgOpenAuction
A new auction is opened using the `MsgOpenAuction` message.

```go
type MsgOpenAuction struct {
	AuctionType	uint32
	NftDenomId	string
	NftTokenId	string
	Duration	uint64
	MinAmount	sdk.Coin
	Owner 		string
	Limit		uint32
}
```

## MsgEditAuction
The `Duration` of an auction can be edited using the `MsgEditAuction` message.

```go
type MsgEditAuction struct {
	Id			uint64	// Auction Id
	Duration	uint64
	Owner		string
}
```

## MsgCancelAuction
An auction can be cancelled using the `MsgCancelAuction` message.

```go
type MsgCancelAuction struct {
	Id			uint64	// Auction Id
	Owner		string
}
```

## MsgOpenBid
A bid can be opened using the `MsgOpenBid` message.

```go
type MsgOpenBid struct {
	AuctionId	uint64
	Bidder		string
	BidAmount	sdk.Coin
}
```

## MsgCancelBid
A bid can be cancelled using the `MsgCancelBid` message.

```go
type MsgCancelBid struct {
	AuctionId	uint64
	Bidder		string
}
```

## MsgWithdraw
The winner can get the NFT and the auction owner can withdraw the tokens through `MsgWithdraw` after the auction is ended.

```go
type MsgWithdraw struct {
	AuctionId	uint64
	Recipient	string
}
```
