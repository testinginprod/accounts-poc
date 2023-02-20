package allowance

import (
	v1 "accounts/examples/allowance/v1"
	"accounts/sdk"
	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	"fmt"
)

func NewAccount(sb *collections.SchemaBuilder) Account {
	return Account{
		Owner:      collections.NewItem(sb, collections.NewPrefix(0), "owner", sdk.IdentityValue),
		Allowances: collections.NewMap(sb, collections.NewPrefix(1), "allowances", collections.PairKeyCodec(sdk.IdentityKey, collections.StringKey), sdk.IntValue),
	}
}

type Account struct {
	Owner      collections.Item[sdk.Identity]
	Allowances collections.Map[collections.Pair[sdk.Identity, string], math.Int]
}

func (a Account) Init(ctx *sdk.Context, initMsg v1.InitMsg) (*sdk.InitResponse, error) {
	return new(sdk.InitResponse), nil
}

func (a Account) IncreaseAllowance(ctx *sdk.Context, user sdk.Identity, denom string, amount math.Int) error {
	owner, err := a.Owner.Get(ctx)
	if err != nil {
		return err
	}
	if !ctx.Sender.Equals(owner) {
		return fmt.Errorf("not authorized")
	}

	allowance, err := a.Allowances.Get(ctx, collections.Join(user, denom))
	if err != nil {
		allowance = math.ZeroInt()
	}
	allowance = allowance.Add(amount)
	return a.Allowances.Set(ctx, collections.Join(user, denom), allowance)
}

func (a Account) RegisterExecuteHandler(r *sdk.ExecuteRouter) {
	sdk.RegisterExecuteHandler(r, func(ctx *sdk.Context, msg v1.IncreaseAllowance) (*sdk.ExecuteResponse, error) {
		panic("impl")
	})

	sdk.RegisterExecuteHandler(r, func(ctx *sdk.Context, msg v1.DecreaseAllowance) (*sdk.ExecuteResponse, error) {
		panic("impl")
	})

	sdk.RegisterExecuteHandler(r, func(ctx *sdk.Context, msg v1.ResetAllowance) (*sdk.ExecuteResponse, error) {
		panic("impl")
	})
}

func (a Account) RegisterQueryHandler(r *sdk.QueryRouter) {

}
