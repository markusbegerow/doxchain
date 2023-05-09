package keeper_test

import (
	"context"
	"testing"

	keepertest "doxchain/testutil/keeper"
	"doxchain/x/did/keeper"
	"doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DidKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
