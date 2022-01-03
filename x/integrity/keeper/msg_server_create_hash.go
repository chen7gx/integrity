package keeper

import (
	"context"

	"github.com/chen7gx/integrity/x/integrity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateHash(goCtx context.Context, msg *types.MsgCreateHash) (*types.MsgCreateHashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message

	var datahash = types.Datahash{
		Creator: msg.Creator,
		Details: msg.Details,
		Hash:    msg.Hash,
	}
	// Add a post to the store and get back the ID
	id := k.AppendDatahash(ctx, datahash)

	// Return the ID of the post

	txgas := sdk.NewInt(1)
	coin := sdk.NewCoin("stake", txgas)
	coins := sdk.NewCoins(coin)

	err := k.MintCoinsForHash(ctx, coins)
	if err != nil {
		panic("mintcoins err")
	}

	creatorAddr, err2 := sdk.AccAddressFromBech32(datahash.Creator)
	if err2 != nil {
		panic("address err")
	}
	err3 := k.SendCoinsFromMintModuleToAccount(ctx, coins, creatorAddr)
	if err3 != nil {
		panic("sendta err")
	}

	return &types.MsgCreateHashResponse{Id: id}, nil
}
