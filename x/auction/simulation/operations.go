package simulation

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/bitsongofficial/chainmodules/types"
	"github.com/bitsongofficial/chainmodules/x/auction/keeper"
	auctiontypes "github.com/bitsongofficial/chainmodules/x/auction/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgOpenAuction   = "op_weight_msg_open_auction"
	OpWeightMsgEditAuction   = "op_weight_msg_edit_auction"
	OpWeightMsgCancelAuction = "op_weight_msg_cancel_auction"
	OpWeightMsgOpenBid       = "op_weight_msg_open_bid"
	OpWeightMsgCancelBid     = "op_weight_msg_cancel_bid"
	OpWeightMsgWithdraw      = "op_weight_msg_withdraw"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc codec.JSONMarshaler,
	k keeper.Keeper,
	ak auctiontypes.AccountKeeper,
	bk auctiontypes.BankKeeper,
) simulation.WeightedOperations {

	var weightOpenAuction, weightEditAuction, weightCancelAuction, weightOpenBid, weightCancelBid int
	appParams.GetOrGenerate(
		cdc, OpWeightMsgOpenAuction, &weightOpenAuction, nil,
		func(_ *rand.Rand) {
			weightOpenAuction = 50
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgEditAuction, &weightEditAuction, nil,
		func(_ *rand.Rand) {
			weightEditAuction = 50
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgCancelAuction, &weightCancelAuction, nil,
		func(_ *rand.Rand) {
			weightCancelAuction = 50
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgOpenBid, &weightOpenBid, nil,
		func(_ *rand.Rand) {
			weightOpenBid = 50
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgCancelBid, &weightCancelBid, nil,
		func(_ *rand.Rand) {
			weightCancelBid = 50
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightOpenAuction,
			SimulateOpenAuction(k, ak, bk),
		),
		simulation.NewWeightedOperation(
			weightEditAuction,
			SimulateEditAuction(k, ak, bk),
		),
		simulation.NewWeightedOperation(
			weightOpenBid,
			SimulateOpenBid(k, ak, bk),
		),
		simulation.NewWeightedOperation(
			weightCancelBid,
			SimulateCancelBid(k, ak, bk),
		),
		simulation.NewWeightedOperation(
			weightCancelAuction,
			SimulateCancelAuction(k, ak, bk),
		),
	}
}

// SimulateOpenAuction tests and runs a single msg open a new auction
func SimulateOpenAuction(k keeper.Keeper, ak auctiontypes.AccountKeeper, bk auctiontypes.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		auction := genAuction(ctx, r, k, ak, bk, accs)
		msg := auctiontypes.NewMsgOpenAuction(auction.AuctionType, auction.NftDenomId, auction.NftTokenId, auction.Duration, auction.MinAmount, auction.Owner, auction.Limit)

		simAccount, found := simtypes.FindAccount(accs, auction.GetOwner())
		if !found {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), fmt.Sprintf("account %s not found", auction.Owner)), nil, fmt.Errorf("account %s not found", auction.Owner)
		}

		owner, _ := sdk.AccAddressFromBech32(msg.Owner)
		account := ak.GetAccount(ctx, owner)
		fees, err := simtypes.RandomFees(r, ctx, sdk.Coins{sdk.NewCoin(types.BondDenom, sdk.NewInt(1000))})
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err = app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "simulate open auction"), nil, nil
	}
}

// SimulateEditAuction tests and runs a single msg edit an existed auction
func SimulateEditAuction(k keeper.Keeper, ak auctiontypes.AccountKeeper, bk auctiontypes.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		auctions := k.GetAuctions(ctx, nil)
		auction := auctions[0]
		msg := auctiontypes.NewMsgEditAuction(auction.Id, auction.Duration+86400, auction.Owner)

		simAccount, found := simtypes.FindAccount(accs, auction.GetOwner())
		if !found {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), fmt.Sprintf("account %s not found", auction.GetOwner())), nil, fmt.Errorf("account %s not found", auction.GetOwner())
		}

		owner, _ := sdk.AccAddressFromBech32(msg.Owner)
		account := ak.GetAccount(ctx, owner)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err = app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "simulate edit auction"), nil, nil
	}
}

// SimulateCancelAuction tests and runs a single msg cancel an existed auction
func SimulateCancelAuction(k keeper.Keeper, ak auctiontypes.AccountKeeper, bk auctiontypes.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		auctions := k.GetAuctions(ctx, nil)
		auction := auctions[0]
		msg := auctiontypes.NewMsgCancelAuction(auction.Id, auction.Owner)

		simAccount, found := simtypes.FindAccount(accs, auction.GetOwner())
		if !found {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), fmt.Sprintf("account %s not found", auction.GetOwner())), nil, fmt.Errorf("account %s not found", auction.GetOwner())
		}

		owner, _ := sdk.AccAddressFromBech32(msg.Owner)
		account := ak.GetAccount(ctx, owner)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err = app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "simulate cancel auction"), nil, nil
	}
}

// SimulateOpenBid tests and runs a single msg open a new bid
func SimulateOpenBid(k keeper.Keeper, ak auctiontypes.AccountKeeper, bk auctiontypes.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		bid := genBid(ctx, r, k, ak, bk, accs)
		msg := auctiontypes.NewMsgOpenBid(bid.AuctionId, bid.Bidder, bid.BidAmount)

		simAccount, found := simtypes.FindAccount(accs, bid.GetBidder())
		if !found {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), fmt.Sprintf("account %s not found", bid.Bidder)), nil, fmt.Errorf("account %s not found", bid.Bidder)
		}

		owner, _ := sdk.AccAddressFromBech32(msg.Bidder)
		account := ak.GetAccount(ctx, owner)
		fees, err := simtypes.RandomFees(r, ctx, sdk.Coins{sdk.NewCoin(types.BondDenom, sdk.NewInt(1000))})
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err = app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "simulate open bid"), nil, nil
	}
}

// SimulateCancelBid tests and runs a single msg cancel an existed bid
func SimulateCancelBid(k keeper.Keeper, ak auctiontypes.AccountKeeper, bk auctiontypes.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		auctions := k.GetAuctions(ctx, nil)
		auction := auctions[0]
		bids := k.GetBidsByAuctionId(ctx, auction.Id)
		bid := bids[0]
		msg := auctiontypes.NewMsgCancelBid(auction.Id, bid.Bidder)

		simAccount, found := simtypes.FindAccount(accs, bid.GetBidder())
		if !found {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), fmt.Sprintf("account %s not found", bid.GetBidder())), nil, fmt.Errorf("account %s not found", bid.GetBidder())
		}

		owner, _ := sdk.AccAddressFromBech32(msg.Bidder)
		account := ak.GetAccount(ctx, owner)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err = app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(auctiontypes.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "simulate cancel bid"), nil, nil
	}
}

func randStringBetween(r *rand.Rand, min, max int) string {
	strLen := simtypes.RandIntBetween(r, min, max)
	randStr := simtypes.RandStringOfLength(r, strLen)
	return strings.ToLower(randStr)
}

func genAuction(ctx sdk.Context,
	r *rand.Rand,
	k keeper.Keeper,
	ak auctiontypes.AccountKeeper,
	bk auctiontypes.BankKeeper,
	accs []simtypes.Account,
) auctiontypes.Auction {

	var auction auctiontypes.Auction
	auction = randAuction(r, accs)

	for k.HasAuction(ctx, auction.GetId()) {
		auction = randAuction(r, accs)
	}

	simAccount, _ := simtypes.RandomAcc(r, accs)
	account := simAccount.Address
	auction.Owner = account.String()

	return auction
}

func genBid(ctx sdk.Context,
	r *rand.Rand,
	k keeper.Keeper,
	ak auctiontypes.AccountKeeper,
	bk auctiontypes.BankKeeper,
	accs []simtypes.Account,
) auctiontypes.Bid {
	auctions := k.GetAuctions(ctx, nil)
	auction := auctions[0]

	bid := randBid(r, auction.Id, accs)

	simAccount, _ := simtypes.RandomAcc(r, accs)
	account := simAccount.Address
	bid.Bidder = account.String()

	return bid
}

func randAuction(r *rand.Rand, accs []simtypes.Account) auctiontypes.Auction {
	id := simtypes.RandIntBetween(r, 0, 1000000)
	simAccount, _ := simtypes.RandomAcc(r, accs)

	return auctiontypes.Auction{
		Id:          uint64(id),
		AuctionType: 0,
		NftDenomId:  randStringBetween(r, 5, 10),
		NftTokenId:  randStringBetween(r, 5, 10),
		StartTime:   0,
		Duration:    86400,
		MinAmount:   sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)),
		Owner:       simAccount.Address.String(),
		Limit:       1,
		Cancelled:   false,
	}
}

func randBid(r *rand.Rand, auctionId uint64, accs []simtypes.Account) auctiontypes.Bid {
	simAccount, _ := simtypes.RandomAcc(r, accs)

	return auctiontypes.Bid{
		AuctionId: auctionId,
		Bidder:    simAccount.Address.String(),
		BidAmount: sdk.NewCoin(types.BondDenom, sdk.NewInt(1000)),
	}
}
