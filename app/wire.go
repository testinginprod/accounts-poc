package app

import (
	"accounts/examples/allowance"
	allowancev1 "accounts/examples/allowance/v1"
	"accounts/x/accounts/keeper"
)

func Accounts() keeper.Accounts {
	return keeper.Accounts{
		keeper.WithAccount[allowancev1.InitMsg]("allowance", allowance.NewAccount),
	}
}
