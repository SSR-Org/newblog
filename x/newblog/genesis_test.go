package newblog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "newblog/testutil/keeper"
	"newblog/testutil/nullify"
	"newblog/x/newblog"
	"newblog/x/newblog/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NewblogKeeper(t)
	newblog.InitGenesis(ctx, *k, genesisState)
	got := newblog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
