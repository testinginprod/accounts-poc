package keeper

import (
	"accounts/sdk"
	"cosmossdk.io/collections"
	"fmt"
	"github.com/gogo/protobuf/proto"
)

type Accounts = []func(sb *collections.SchemaBuilder) (*InternalAccount, error)

type InternalAccount struct {
	name    func() string
	init    func(ctx *sdk.Context, msg proto.Message) (*sdk.InitResponse, error)
	execute func(ctx *sdk.Context, msg proto.Message) (*sdk.ExecuteResponse, error)
	query   func(ctx *sdk.Context, msg proto.Message) (proto.Message, error)
	schema  func() *schema
}

func WithAccount[IM any, PIM sdk.Encodable[IM], A sdk.Account[IM, PIM]](name string, accCreator func(sb *collections.SchemaBuilder) A) func(sb *collections.SchemaBuilder) (*InternalAccount, error) {
	return func(sb *collections.SchemaBuilder) (*InternalAccount, error) {
		return NewInternalAccount[IM, PIM, A](sb, func(sb *collections.SchemaBuilder) (string, A) {
			return name, accCreator(sb)
		})
	}
}

func NewInternalAccount[IM any, PIM sdk.Encodable[IM], A sdk.Account[IM, PIM]](sb *collections.SchemaBuilder, accCreator func(sb *collections.SchemaBuilder) (string, A)) (*InternalAccount, error) {
	name, account := accCreator(sb)
	stateSchema, err := sb.Build()
	if err != nil {
		return nil, err
	}

	executeHandler := sdk.NewExecuteRouter()
	account.RegisterExecuteHandler(executeHandler)

	queryHandler := sdk.NewQueryRouter()
	account.RegisterQueryHandler(queryHandler)

	return &InternalAccount{
		name: func() string { return name },

		init: func(ctx *sdk.Context, msg proto.Message) (*sdk.InitResponse, error) {
			concrete, ok := msg.(PIM)
			if !ok {
				return nil, fmt.Errorf("invalid init request, wanted %T, got: %T", new(IM), msg)
			}
			return account.Init(ctx, *concrete)
		},
		execute: executeHandler.Handler(),

		query: queryHandler.Handler(),

		schema: func() *schema {
			return &schema{
				messages: nil,
				queries:  nil,
				state:    &stateSchema,
			}
		},
	}, nil
}
