package keeper_test

import (
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/did/keeper"
	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNDid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Did {
	items := make([]types.Did, n)
	for i := range items {
		items[i].Id = keeper.AppendDid(ctx, items[i])
	}
	return items
}

func TestDidGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDid(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetDid(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestDidRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDid(ctx, item.Id)
		_, found := keeper.GetDid(ctx, item.Id)
		require.False(t, found)
	}
}

func TestDidGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDid(ctx)),
	)
}

func TestDidCount(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDid(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetDidCount(ctx))
}
