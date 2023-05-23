package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
)

// GetMaxExpectedTimePerBlock retrieves the maximum expected time per block from the paramstore
func (k Keeper) GetMaxExpectedTimePerBlock(ctx sdk.Context) uint64 {
	p := k.GetParams(ctx)
	return p.MaxExpectedTimePerBlock
}

// GetParams returns the total set of ibc-connection parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.ParamsKey))
	if bz == nil {
		return p
	}

	k.cdc.MustUnmarshal(bz, &p)
	return p
}

// SetParams sets the total set of ibc-connection parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set([]byte(types.ParamsKey), bz)
}
