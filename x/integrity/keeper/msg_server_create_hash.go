package keeper

import (
	"context"

	"github.com/chen7gx/integrity/x/integrity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateHash(goCtx context.Context, msg *types.MsgCreateHash) (*types.MsgCreateHashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateHashResponse{}, nil
}
