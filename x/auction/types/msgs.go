package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// MsgRoute identifies transaction types
	MsgRoute = "auction"

	TypeMsgOpenAuction   = "open_auction"
	TypeMsgEditAuction   = "edit_auction"
	TypeMsgCancelAuction = "cancel_auction"
	TypeMsgOpenBid       = "open_bid"
	TypeMsgCancelBid     = "cancel_bid"
	TypeMsgWithdraw      = "withdraw"

	// DoNotModify used to indicate that some field should not be updated
	DoNotModify = "[do-not-modify]"
)

var (
	_ sdk.Msg = &MsgOpenAuction{}
	_ sdk.Msg = &MsgEditAuction{}
	_ sdk.Msg = &MsgCancelAuction{}
	_ sdk.Msg = &MsgOpenBid{}
	_ sdk.Msg = &MsgCancelBid{}
	_ sdk.Msg = &MsgWithdraw{}
)

// NewMsgOpenAuction - construct auction open msg.
func NewMsgOpenAuction(
	auctionType AuctionType, nftDenomId string, nftTokenId string, duration uint64,
	minAmount sdk.Coin, owner string, limit uint32,
) *MsgOpenAuction {
	return &MsgOpenAuction{
		AuctionType: auctionType,
		NftDenomId:  nftDenomId,
		NftTokenId:  nftTokenId,
		Duration:    duration,
		MinAmount:   minAmount,
		Owner:       owner,
		Limit:       limit,
	}
}

// Route Implements Msg.
func (msg MsgOpenAuction) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgOpenAuction) Type() string { return TypeMsgOpenAuction }

// ValidateBasic Implements Msg.
func (msg MsgOpenAuction) ValidateBasic() error {
	now := time.Now()
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return err
	}

	return ValidateAuction(
		NewAuction(
			1, // auctionIdForValidation
			msg.AuctionType,
			msg.NftDenomId,
			msg.NftTokenId,
			uint64(now.Unix()),
			msg.Duration,
			msg.MinAmount,
			owner,
			msg.Limit,
		),
	)
}

// GetSignBytes Implements Msg.
func (msg MsgOpenAuction) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgOpenAuction) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgEditAuction - construct auction edit msg.
func NewMsgEditAuction(
	id uint64, duration uint64, owner string,
) *MsgEditAuction {
	return &MsgEditAuction{
		Id:       id,
		Duration: duration,
		Owner:    owner,
	}
}

// Route Implements Msg.
func (msg MsgEditAuction) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgEditAuction) Type() string { return TypeMsgEditAuction }

// ValidateBasic Implements Msg.
func (msg MsgEditAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return ValidateDuration(msg.Duration)
}

// GetSignBytes Implements Msg.
func (msg MsgEditAuction) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgEditAuction) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgCancelAuction - construct auction cancel msg.
func NewMsgCancelAuction(
	id uint64, owner string,
) *MsgCancelAuction {
	return &MsgCancelAuction{
		Id:    id,
		Owner: owner,
	}
}

// Route Implements Msg.
func (msg MsgCancelAuction) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgCancelAuction) Type() string { return TypeMsgCancelAuction }

// ValidateBasic Implements Msg.
func (msg MsgCancelAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCancelAuction) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgCancelAuction) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgOpenBid - construct bid open msg.
func NewMsgOpenBid(
	auctionId uint64, bidder string, bidAmount sdk.Coin,
) *MsgOpenBid {
	return &MsgOpenBid{
		AuctionId: auctionId,
		Bidder:    bidder,
		BidAmount: bidAmount,
	}
}

// Route Implements Msg.
func (msg MsgOpenBid) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgOpenBid) Type() string { return TypeMsgOpenAuction }

// ValidateBasic Implements Msg.
func (msg MsgOpenBid) ValidateBasic() error {
	bidder, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return err
	}
	return ValidateBid(
		NewBid(
			msg.AuctionId,
			bidder,
			msg.BidAmount,
		),
	)
}

// GetSignBytes Implements Msg.
func (msg MsgOpenBid) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgOpenBid) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgCancelBid - construct bid cancel msg.
func NewMsgCancelBid(
	auctionId uint64, bidder string,
) *MsgCancelBid {
	return &MsgCancelBid{
		AuctionId: auctionId,
		Bidder:    bidder,
	}
}

// Route Implements Msg.
func (msg MsgCancelBid) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgCancelBid) Type() string { return TypeMsgCancelBid }

// ValidateBasic Implements Msg.
func (msg MsgCancelBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCancelBid) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgCancelBid) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgCancelAuction - construct auction cancel msg.
func NewMsgWithdraw(
	auctionId uint64, recipient string,
) *MsgWithdraw {
	return &MsgWithdraw{
		AuctionId: auctionId,
		Recipient: recipient,
	}
}

// Route Implements Msg.
func (msg MsgWithdraw) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgWithdraw) Type() string { return TypeMsgWithdraw }

// ValidateBasic Implements Msg.
func (msg MsgWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgWithdraw) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
