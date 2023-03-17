package recover

import (
	v1 "accounts/examples/recover/v1"
	"accounts/sdk"
	authnv1 "accounts/x/authn/types"
	"cosmossdk.io/collections"
	"fmt"
)

func NewAccount(sb *collections.SchemaBuilder) Account {
	return Account{
		Recoverer: collections.NewItem(sb, collections.NewPrefix(0), "recoverer", sdk.AccAddressValue),
	}
}

type Account struct {
	Recoverer collections.Item[sdk.AccAddress]
}

func (a Account) Init(ctx *sdk.Context, initMsg v1.Init) (*sdk.InitResponse, error) {
	recoverer, err := sdk.AccAddressFromBech32(initMsg.AccountRecoverer)
	if err != nil {
		return nil, err
	}
	err = a.Recoverer.Set(ctx, recoverer)
	if err != nil {
		return nil, err
	}
	return new(sdk.InitResponse), nil
}

func (a Account) ChangePubKey(ctx *sdk.Context, newCredentials []byte) (*sdk.ExecuteResponse, error) {
	recoverer, err := a.Recoverer.Get(ctx)
	if err != nil {
		return nil, err
	}

	if !recoverer.Equals(ctx.Sender) {
		return nil, fmt.Errorf("unauthorized")
	}

	return new(sdk.ExecuteResponse).
		WithCosmoSDKMsg(&authnv1.MsgUpdateCredentials{
			Sender:   ctx.Self.String(),
			NewAuthn: newCredentials,
		}), nil
}

func (a Account) RegisterExecuteHandler(router *sdk.ExecuteRouter) {
	sdk.RegisterExecuteHandler(router, func(ctx *sdk.Context, msg v1.MsgSwapPublicKey) (*sdk.ExecuteResponse, error) {
		return a.ChangePubKey(ctx, msg.NewPublicKey)
	})
}

func (a Account) RegisterQueryHandler(router *sdk.QueryRouter) {}
