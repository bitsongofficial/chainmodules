package types

import (
	"github.com/bitsongofficial/chainmodules/x/nft/exported"
	nfttypes "github.com/bitsongofficial/chainmodules/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// BankKeeper defines the expected bank keeper (noalias)
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAddress(name string) sdk.AccAddress
}

type NftKeeper interface {
	TransferOwner(
		ctx sdk.Context, denomID, tokenID string, srcOwner, dstOwner sdk.AccAddress,
	) error
	GetDenom(ctx sdk.Context, id string) (denom nfttypes.Denom, err error)
	GetNFT(ctx sdk.Context, denomID, tokenID string) (nft exported.NFT, err error)
	SetNFT(ctx sdk.Context, denomID string, nft nfttypes.BaseNFT)
	CloneMintNFT(
		ctx sdk.Context, denomID, tokenID, tokenNm,
		tokenURI string, owner sdk.AccAddress,
	) error
}
