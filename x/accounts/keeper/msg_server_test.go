package keeper_test

import (
	"context"
	"testing"

	keepertest "accounts/testutil/keeper"
	"accounts/x/accounts/keeper"
	"accounts/x/accounts/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AccountsKeeper(t, nil)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
