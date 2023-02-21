package keeper

import "cosmossdk.io/collections"

type schema struct {
	messages interface{}
	queries  interface{}
	state    *collections.Schema
}

type runtime struct {
	accounts map[string]InternalAccount
	schemas  map[string]*schema
}
