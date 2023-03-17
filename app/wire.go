package app

import (
	"accounts/examples/recover"
	recoverv1 "accounts/examples/recover/v1"
	"accounts/examples/vesting"
	vestingv1 "accounts/examples/vesting/v1"

	"accounts/x/accounts/keeper"
)

func Accounts() keeper.Accounts {
	return keeper.Accounts{
		keeper.WithAccount[vestingv1.Init]("simple-vesting", vesting.NewAccount),
		keeper.WithAccount[recoverv1.Init]("recoverable-account", recover.NewAccount),
	}
}
