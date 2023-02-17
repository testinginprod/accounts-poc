package keeper

import (
	"accounts/x/accounts/types"
)

var _ types.QueryServer = Keeper{}
