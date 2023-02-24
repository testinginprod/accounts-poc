package keeper_test

import (
	"accounts/app"
	vestingv1 "accounts/examples/vesting/v1"
	keepertest "accounts/testutil/keeper"
	"accounts/testutil/sample"
	"accounts/utils"
	"accounts/x/accounts/keeper"
	accountstypes "accounts/x/accounts/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	k, ctx := keepertest.AccountsKeeper(t, app.Accounts())
	s := keeper.NewMsgServerImpl(*k)

	r, err := s.Deploy(sdk.WrapSDKContext(ctx), &accountstypes.MsgDeploy{
		Sender: sample.AccAddress(),
		Kind:   "simple-vesting",
		InitMessage: anyfy(t, &vestingv1.Init{
			Beneficiary: sample.AccAddress(),
			StartAfter:  10 * time.Second,
			Duration:    100 * time.Second,
		}),
		Funds: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1000))),
	})

	require.NoError(t, err)
	t.Logf("%s", r)
}

func anyfy(t *testing.T, msg proto.Message) []byte {
	// {"beneficiary":"someone","startAfter":"10s","duration":"100s"}
	b, err := utils.MarshalAnyBytes(msg)
	require.NoError(t, err)
	return b
}
