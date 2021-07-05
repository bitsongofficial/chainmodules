//nolint
package keeper

import (
	"math"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitsongofficial/go-bitsong/x/fantoken/types"
)

// fee factor formula: (ln(len({name}))/ln{base})^{exp}
const (
	FeeFactorBase = 3
	FeeFactorExp  = 4
)

// DeductIssueTokenFee performs fee handling for issuing token
func (k Keeper) DeductIssueFanTokenFee(ctx sdk.Context, owner sdk.AccAddress, issueFee sdk.Coin, symbol string) error {
	// get the required issuance fee
	fee := k.GetFanTokenIssueFee(ctx, issueFee, symbol)
	return feeHandler(ctx, k, owner, fee)
}

// GetTokenIssueFee returns the token issuance fee
func (k Keeper) GetFanTokenIssueFee(ctx sdk.Context, issueFee sdk.Coin, symbol string) sdk.Coin {
	fee := k.calcFanTokenIssueFee(ctx, issueFee, symbol)
	return fee
}

func (k Keeper) calcFanTokenIssueFee(ctx sdk.Context, issueFee sdk.Coin, symbol string) sdk.Coin {
	// compute the fee
	feeAmt := calcFeeByBase(symbol, issueFee.Amount)
	if feeAmt.GT(sdk.NewDec(1)) {
		return sdk.NewCoin(issueFee.Denom, feeAmt.TruncateInt())
	}
	return sdk.NewCoin(issueFee.Denom, sdk.OneInt())
}

// feeHandler handles the fee of token
func feeHandler(ctx sdk.Context, k Keeper, feeAcc sdk.AccAddress, fee sdk.Coin) error {
	burnedCoins := sdk.NewCoins(fee)

	// send all fees to module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, feeAcc, types.ModuleName, sdk.NewCoins(fee),
	); err != nil {
		return err
	}

	// burn burnedCoin
	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnedCoins)
}

// calcFeeByBase computes the actual fee according to the given base fee
func calcFeeByBase(name string, baseFee sdk.Int) sdk.Dec {
	feeFactor := calcFeeFactor(name)
	actualFee := sdk.NewDecFromInt(baseFee).Quo(feeFactor)

	return actualFee
}

// calcFeeFactor computes the fee factor of the given name
// Note: make sure that the name size is examined before invoking the function
func calcFeeFactor(name string) sdk.Dec {
	nameLen := len(name)
	if nameLen == 0 {
		panic("the length of name must be greater than 0")
	}

	denominator := math.Log(FeeFactorBase)
	numerator := math.Log(float64(nameLen))

	feeFactor := math.Pow(numerator/denominator, FeeFactorExp)
	feeFactorDec, err := sdk.NewDecFromStr(strconv.FormatFloat(feeFactor, 'f', 2, 64))
	if err != nil {
		panic("invalid string")
	}

	return feeFactorDec
}
