package keeper

import (
	"accounts/sdk"
	"cosmossdk.io/collections"
	"fmt"
	"github.com/gogo/protobuf/proto"
)

type Accounts = []func(sb *collections.SchemaBuilder) InternalAccount

type InternalAccount interface {
	init(ctx *sdk.Context, msg proto.Message) (*sdk.InitResponse, error)
	execute(ctx *sdk.Context, msg proto.Message) (*sdk.ExecuteResponse, error)
	query(ctx *sdk.Context, msg proto.Message) (proto.Message, error)
}

type internalAccountImpl[IM any, PIM sdk.Encodable[IM]] struct {
	initter  func(ctx *sdk.Context, initMsg IM) (*sdk.InitResponse, error)
	executor func(ctx *sdk.Context, msg proto.Message) (*sdk.ExecuteResponse, error)
	querier  func(ctx *sdk.Context, msg proto.Message) (proto.Message, error)
}

func (a *internalAccountImpl[IM, PIM]) init(ctx *sdk.Context, msg proto.Message) (*sdk.InitResponse, error) {
	concrete, ok := msg.(PIM)
	if !ok {
		return nil, fmt.Errorf("invalid request")
	}
	return a.initter(ctx, *concrete)
}

func (a *internalAccountImpl[IM, PIM]) execute(ctx *sdk.Context, msg proto.Message) (*sdk.ExecuteResponse, error) {
	return a.executor(ctx, msg)
}

func (a *internalAccountImpl[IM, PIM]) query(ctx *sdk.Context, msg proto.Message) (proto.Message, error) {
	return a.querier(ctx, msg)
}

func WithAccount[IM any, PIM sdk.Encodable[IM], A sdk.Account[IM, PIM]](accCreator func(sb *collections.SchemaBuilder) A) func(sb *collections.SchemaBuilder) InternalAccount {
	return func(sb *collections.SchemaBuilder) InternalAccount {
		account := accCreator(sb)
		er := sdk.NewExecuteRouter()
		qr := sdk.NewQueryRouter()

		account.RegisterExecuteHandler(er)
		account.RegisterQueryHandler(qr)

		return &internalAccountImpl[IM, PIM]{
			initter:  account.Init,
			executor: er.Handler(),
			querier:  qr.Handler(),
		}
	}
}
