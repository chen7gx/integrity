package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/chen7gx/integrity/testutil/keeper"
	"github.com/chen7gx/integrity/x/integrity/keeper"
	"github.com/chen7gx/integrity/x/integrity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IntegrityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
