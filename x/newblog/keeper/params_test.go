package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "newblog/testutil/keeper"
	"newblog/x/newblog/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NewblogKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
