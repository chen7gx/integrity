package integrity_test

import (
	"testing"

	keepertest "github.com/chen7gx/integrity/testutil/keeper"
	"github.com/chen7gx/integrity/x/integrity"
	"github.com/chen7gx/integrity/x/integrity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IntegrityKeeper(t)
	integrity.InitGenesis(ctx, *k, genesisState)
	got := integrity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
