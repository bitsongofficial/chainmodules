package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

func (k Keeper) getLastAuctionId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.KeyLastAuctionId)
	if bytes == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bytes)
}

func (k Keeper) setLastAuctionId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLastAuctionId, sdk.Uint64ToBigEndian(id))
}
