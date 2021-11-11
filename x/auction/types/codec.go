package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*AuctionI)(nil), nil)
	cdc.RegisterInterface((*BidI)(nil), nil)

	cdc.RegisterConcrete(&Auction{}, "go-bitsong/auction/Auction", nil)
	cdc.RegisterConcrete(&Bid{}, "go-bitsong/auction/Bid", nil)

	cdc.RegisterConcrete(&MsgOpenAuction{}, "go-bitsong/auction/MsgOpenAuction", nil)
	cdc.RegisterConcrete(&MsgEditAuction{}, "go-bitsong/auction/MsgEditAuction", nil)
	cdc.RegisterConcrete(&MsgCancelAuction{}, "go-bitsong/auction/MsgCancelAuction", nil)
	cdc.RegisterConcrete(&MsgOpenBid{}, "go-bitsong/auction/MsgOpenBid", nil)
	cdc.RegisterConcrete(&MsgCancelBid{}, "go-bitsong/auction/MsgCancelBid", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "go-bitsong/auction/MsgWithdraw", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOpenAuction{},
		&MsgEditAuction{},
		&MsgCancelAuction{},
		&MsgOpenBid{},
		&MsgCancelBid{},
		&MsgWithdraw{},
	)
	registry.RegisterInterface(
		"go-bitsong.auction.AuctionI",
		(*AuctionI)(nil),
		&Auction{},
	)
	registry.RegisterInterface(
		"go-bitsong.auction.BidI",
		(*BidI)(nil),
		&Bid{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
