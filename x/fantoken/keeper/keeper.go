package keeper

import (
	"fmt"
	"strings"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/bitsongofficial/chainmodules/x/fantoken/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
)

type Keeper struct {
	storeKey         sdk.StoreKey
	cdc              codec.Marshaler
	bankKeeper       types.BankKeeper
	paramSpace       paramstypes.Subspace
	blockedAddrs     map[string]bool
	feeCollectorName string
}

func NewKeeper(
	cdc codec.Marshaler,
	key sdk.StoreKey,
	paramSpace paramstypes.Subspace,
	bankKeeper types.BankKeeper,
	blockedAddrs map[string]bool,
	feeCollectorName string,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:         key,
		cdc:              cdc,
		paramSpace:       paramSpace,
		bankKeeper:       bankKeeper,
		feeCollectorName: feeCollectorName,
		blockedAddrs:     blockedAddrs,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("go-bitsong/%s", types.ModuleName))
}

// IssueFanToken issues a new fantoken
func (k Keeper) IssueFanToken(
	ctx sdk.Context,
	symbol string,
	name string,
	maxSupply sdk.Int,
	description string,
	owner sdk.AccAddress,
	issueFee sdk.Coin,
) error {
	issuePrice := k.GetParamSet(ctx).IssuePrice
	if issueFee.Denom != issuePrice.GetDenom() {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "the issue fee denom %s is invalid", issueFee.String())
	}
	if issueFee.Amount.LT(issuePrice.Amount) {
		return sdkerrors.Wrapf(types.ErrLessIssueFee, "the issue fee %s is less than %s", issueFee.String(), issuePrice.String())
	}

	denom := strings.Replace(common.BytesToHash([]byte(owner.String()+symbol+name)).Hex(), "0x", "ft", 1)
	denomMetaData := banktypes.Metadata{
		Description: description,
		Base:        denom,
		Display:     symbol,
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: denom, Exponent: 0},
			{Denom: symbol, Exponent: types.FanTokenDecimal},
		},
	}
	token := types.NewFanToken(name, maxSupply, owner, denomMetaData)

	if err := k.AddFanToken(ctx, token); err != nil {
		return err
	}

	return nil
}

// EditFanToken edits the specified fantoken
func (k Keeper) EditFanToken(
	ctx sdk.Context,
	denom string,
	mintable bool,
	owner sdk.AccAddress,
) error {
	// get the destination token
	token, err := k.getFanTokenByDenom(ctx, denom)
	if err != nil {
		return err
	}

	if owner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", owner, denom)
	}

	if !token.Mintable {
		return sdkerrors.Wrapf(types.ErrNotMintable, "the fantoken %s is not mintable", denom)
	}

	token.Mintable = mintable

	if !mintable {
		supply := k.getFanTokenSupply(ctx, token.GetDenom())
		precision := sdk.NewIntWithDecimal(1, types.FanTokenDecimal)
		token.MaxSupply = supply.Quo(precision)
	}

	k.setFanToken(ctx, token)

	return nil
}

// TransferFanTokenOwner transfers the owner of the specified fantoken to a new one
func (k Keeper) TransferFanTokenOwner(
	ctx sdk.Context,
	denom string,
	srcOwner sdk.AccAddress,
	dstOwner sdk.AccAddress,
) error {
	token, err := k.getFanTokenByDenom(ctx, denom)
	if err != nil {
		return err
	}

	if srcOwner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", srcOwner, denom)
	}

	token.Owner = dstOwner.String()

	// update token
	k.setFanToken(ctx, token)

	// reset all indices
	k.resetStoreKeyForQueryToken(ctx, token.GetDenom(), srcOwner, dstOwner)

	return nil
}

// MintFanToken mints the specified amount of fantoken to the specified recipient
func (k Keeper) MintFanToken(
	ctx sdk.Context,
	recipient sdk.AccAddress,
	denom string,
	amount sdk.Int,
	owner sdk.AccAddress,
) error {
	token, err := k.getFanTokenByDenom(ctx, denom)
	if err != nil {
		return err
	}

	if owner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", owner, denom)
	}

	if !token.Mintable {
		return sdkerrors.Wrapf(types.ErrNotMintable, "%s", denom)
	}

	supply := k.getFanTokenSupply(ctx, token.GetDenom())
	mintableAmt := token.MaxSupply.Sub(supply)

	if amount.GT(mintableAmt) {
		return sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			"the amount exceeds the mintable token amount; expected (0, %d], got %d",
			mintableAmt, amount,
		)
	}

	mintCoin := sdk.NewCoin(token.GetDenom(), amount)
	mintCoins := sdk.NewCoins(mintCoin)

	// mint coins
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return err
	}

	if recipient.Empty() {
		recipient = owner
	}

	// sent coins to the recipient account
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, mintCoins)
}

// BurnToken burns the specified amount of token
func (k Keeper) BurnFanToken(
	ctx sdk.Context,
	denom string,
	amount sdk.Int,
	owner sdk.AccAddress,
) error {
	_, err := k.getFanTokenByDenom(ctx, denom)
	if err != nil {
		return err
	}

	burnCoin := sdk.NewCoin(denom, amount)
	burnCoins := sdk.NewCoins(burnCoin)

	// burn coins
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, burnCoins); err != nil {
		return err
	}

	k.AddBurnCoin(ctx, burnCoin)

	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
}
