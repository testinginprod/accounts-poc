package keeper_test

import (
	"context"
	"testing"

	keepertest "accounts/testutil/keeper"
	"accounts/x/authn/keeper"
	"accounts/x/authn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AuthnKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
