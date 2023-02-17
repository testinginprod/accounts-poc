package accounts_test

import (
	"testing"

	keepertest "accounts/testutil/keeper"
	"accounts/testutil/nullify"
	"accounts/x/accounts"
	"accounts/x/accounts/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AccountsKeeper(t)
	accounts.InitGenesis(ctx, *k, genesisState)
	got := accounts.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
