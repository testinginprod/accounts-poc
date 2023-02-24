package app

import (
	"accounts/examples/vesting"
	vestingv1 "accounts/examples/vesting/v1"
	"accounts/x/accounts/keeper"
)

func Accounts() keeper.Accounts {
	return keeper.Accounts{
		keeper.WithAccount[vestingv1.Init]("simple-vesting", vesting.NewAccount),
	}
}
