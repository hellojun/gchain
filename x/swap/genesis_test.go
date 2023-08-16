package swap_test

import (
	"testing"

	keepertest "gchain/testutil/keeper"
	"gchain/testutil/nullify"
	"gchain/x/swap"
	"gchain/x/swap/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SwapKeeper(t)
	swap.InitGenesis(ctx, *k, genesisState)
	got := swap.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
