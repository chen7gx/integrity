package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/chen7gx/integrity/x/integrity/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		bankKeeper types.BankKeeper
		mintKeeper types.MintKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	bk types.BankKeeper,
	mk types.MintKeeper,

) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		bankKeeper: bk,
		mintKeeper: mk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) MintCoinsForHash(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.mintKeeper.MintCoins(ctx, newCoins)
}

func (k Keeper) SendCoinsFromMintModuleToAccount(ctx sdk.Context, amt sdk.Coins, addr sdk.AccAddress) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "mint", addr, amt)
}
