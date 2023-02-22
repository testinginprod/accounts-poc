package sdk

import (
	"context"
	"cosmossdk.io/core/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type Context struct {
	context.Context // ugly hack
	sdkCtx          *sdk.Context
	Store           store.KVStore
	Sender          Identity
	Funds           Coins
	Self            Identity
	SelfID          uint64
}

func (c *Context) BlockTime() time.Time { return c.sdkCtx.BlockTime() }

func NewContextFromSDK(ctx sdk.Context, sender Identity, self Identity, accountID uint64, store store.KVStore, funds Coins) *Context {
	return &Context{
		Context: context.Background(),
		sdkCtx:  &ctx,
		Store:   store,
		Sender:  sender,
		Funds:   nil,
		Self:    self,
		SelfID:  accountID,
	}
}
