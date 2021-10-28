package keeper

import (
	"encoding/binary"
	"github.com/chen7gx/integrity/x/integrity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendDatahash(ctx sdk.Context, datahash types.Datahash) uint64 {
	// Get the current number of hash in the store
	count := k.GethashCount(ctx)
	// Assign an ID to the hash based on the number of hash in the store
	datahash.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.HashKey))
	// Convert the hash ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, datahash.Id)
	// Marshal the hash into bytes
	appendedValue := k.cdc.MustMarshal(&datahash)
	// Insert the hash bytes using hash ID as a key
	store.Set(byteKey, appendedValue)
	// Update the hash count
	k.SethashCount(ctx, count+1)
	return count
}

func (k Keeper) GethashCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and hashCountKey (which is "hash-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.HashCountKey))
	// Convert the hashCountKey to bytes
	byteKey := []byte(types.HashCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first hash)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SethashCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and hashCountKey (which is "hash-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.HashCountKey))
	// Convert the hashCountKey to bytes
	byteKey := []byte(types.HashCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of hash-count- to count
	store.Set(byteKey, bz)
}
