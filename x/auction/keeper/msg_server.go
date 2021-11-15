package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the token MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) OpenAuction(goCtx context.Context, msg *types.MsgOpenAuction) (*types.MsgOpenAuctionResponse, error) {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	auctionId, err := m.Keeper.OpenAuction(
		ctx, msg.AuctionType, msg.NftDenomId, msg.NftTokenId, msg.Duration, msg.MinAmount, owner, msg.Limit,
	)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeOpenAuction,
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(auctionId)),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgOpenAuctionResponse{
		AuctionId: auctionId,
	}, nil
}

func (m msgServer) EditAuction(goCtx context.Context, msg *types.MsgEditAuction) (*types.MsgEditAuctionResponse, error) {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.EditAuction(
		ctx, msg.Id, msg.Duration, owner,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditAuction,
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(msg.Id)),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgEditAuctionResponse{}, nil
}

func (m msgServer) CancelAuction(goCtx context.Context, msg *types.MsgCancelAuction) (*types.MsgCancelAuctionResponse, error) {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.CancelAuction(
		ctx, msg.Id, owner,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelAuction,
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(msg.Id)),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgCancelAuctionResponse{}, nil
}

func (m msgServer) OpenBid(goCtx context.Context, msg *types.MsgOpenBid) (*types.MsgOpenBidResponse, error) {
	bidder, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.OpenBid(
		ctx, msg.AuctionId, bidder, msg.BidAmount,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeOpenBid,
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(msg.AuctionId)),
			sdk.NewAttribute(types.AttributeKeyBidder, msg.Bidder),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Bidder),
		),
	})

	return &types.MsgOpenBidResponse{}, nil
}

func (m msgServer) CancelBid(goCtx context.Context, msg *types.MsgCancelBid) (*types.MsgCancelBidResponse, error) {
	bidder, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.CancelBid(
		ctx, msg.AuctionId, bidder,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelBid,
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(msg.AuctionId)),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Bidder),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Bidder),
		),
	})

	return &types.MsgCancelBidResponse{}, nil
}

func (m msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.Withdraw(
		ctx, msg.AuctionId, recipient,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdraw,
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(msg.AuctionId)),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Recipient),
		),
	})

	return &types.MsgWithdrawResponse{}, nil
}
