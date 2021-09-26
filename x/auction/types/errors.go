//nolint
package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// auction module sentinel errors
var (
	ErrAuctionNotExists    = sdkerrors.Register(ModuleName, 2, "auction does not exist")
	ErrInvalidOwner        = sdkerrors.Register(ModuleName, 3, "invalid auction owner")
	ErrInvalidAmount       = sdkerrors.Register(ModuleName, 4, "invalid amount")
	ErrInvalidAuctionType  = sdkerrors.Register(ModuleName, 5, "invalid auction type")
	ErrInvalidAuctionLimit = sdkerrors.Register(ModuleName, 6, "invalid auction limit")
	ErrInvalidDuration     = sdkerrors.Register(ModuleName, 7, "invalid duration")
	ErrInvalidAuction      = sdkerrors.Register(ModuleName, 8, "invalid auction")
)
