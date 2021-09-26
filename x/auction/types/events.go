// nolint
package types

const (
	EventTypeOpenAuction   = "open_auction"
	EventTypeEditAuction   = "edit_auction"
	EventTypeCancelAuction = "cancel_auction"
	EventTypeOpenBid       = "open_bid"
	EventTypeCancelBid     = "cancel_bid"
	EventTypeWithdraw      = "withdraw"

	AttributeValueCategory = ModuleName

	AttributeKeyAuctionId = "auction_id"
	AttributeKeyCreator   = "creator"
	AttributeKeyBidder    = "bidder"
	AttributeKeyOwner     = "owner"
	AttributeKeyRecipient = "recipient"
)
