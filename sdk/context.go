package sdk

import (
	"context"
	"cosmossdk.io/core/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type Context struct {
	context.Context // needs to implement ctx, for collections.
	sdkCtx          *sdk.Context
	Store           store.KVStore
	Sender          AccAddress
	Funds           Coins
	Self            AccAddress
	SelfID          uint64
}

type Attribute = sdk.Attribute

func (c *Context) BlockTime() time.Time { return c.sdkCtx.BlockTime() }
func (c *Context) WithEvent(name string, attrs ...Attribute) {
	c.sdkCtx.EventManager().EmitEvent(sdk.NewEvent(name, attrs...))
}

func NewContextFromSDK(ctx sdk.Context, sender AccAddress, self AccAddress, accountID uint64, store store.KVStore, funds Coins) *Context {
	return &Context{
		Context: context.Background(),
		sdkCtx:  &ctx,
		Store:   store,
		Sender:  sender,
		Funds:   funds,
		Self:    self,
		SelfID:  accountID,
	}
}
