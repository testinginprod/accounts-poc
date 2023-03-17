package keeper

import (
	"accounts/x/authn/types"
)

var _ types.QueryServer = Keeper{}
