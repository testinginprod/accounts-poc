package vesting

import (
	v1 "accounts/examples/vesting/v1"
	"accounts/sdk"
	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	"fmt"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"time"
)

func NewAccount(sb *collections.SchemaBuilder) Account {
	return Account{
		Beneficiary:      collections.NewItem(sb, collections.NewPrefix(0), "owner", sdk.AccAddressValue),
		VestedCoinDenom:  collections.NewItem(sb, collections.NewPrefix(1), "vested_denom", collections.StringValue),
		VestedAmount:     collections.NewItem(sb, collections.NewPrefix(2), "vested_amount", sdk.IntValue),
		StartTime:        collections.NewItem(sb, collections.NewPrefix(3), "start_time", sdk.TimeValue),
		UnlocksPerSecond: collections.NewItem(sb, collections.NewPrefix(4), "unlocks_per_second", sdk.IntValue),
		WithdrawnAmount:  collections.NewItem(sb, collections.NewPrefix(5), "withdrawn_amount", sdk.IntValue),
	}
}

type Account struct {
	Beneficiary      collections.Item[sdk.AccAddress] // Beneficiary is the holder of the vested coins.
	VestedCoinDenom  collections.Item[string]         // VestedCoinDenom is the denom of the vested coin.
	VestedAmount     collections.Item[math.Int]       // VestedAmount keeps track of the amount of coins that are vested
	StartTime        collections.Item[time.Time]      // StartTime defines when the account can start withdrawing vested coins.
	UnlocksPerSecond collections.Item[math.Int]       // UnlocksPerSecond defines how many coins unlock after each second.
	WithdrawnAmount  collections.Item[math.Int]       // WithdrawnAmount keeps track of the amount of coins withdrawn so far.
}

func (a Account) Init(ctx *sdk.Context, msg v1.Init) (*sdk.InitResponse, error) {
	// check funds
	if len(ctx.Funds) != 1 {
		return nil, fmt.Errorf("only one coin per vested account, got: %s", ctx.Funds)
	}
	vestedCoin := ctx.Funds[0]
	if msg.Duration <= 0 {
		return nil, fmt.Errorf("invalid duration")
	}

	// check time
	err := a.StartTime.Set(ctx, ctx.BlockTime().Add(msg.StartAfter))
	if err != nil {
		return nil, err
	}

	// check beneficiary
	beneficiary, err := sdk.AccAddressFromBech32(msg.Beneficiary)
	if err != nil {
		return nil, err
	}

	err = a.Beneficiary.Set(ctx, beneficiary)
	if err != nil {
		return nil, err
	}

	unlocksPerSecond := vestedCoin.Amount.QuoRaw(msg.Duration.Milliseconds() / 1000)
	err = a.UnlocksPerSecond.Set(ctx, unlocksPerSecond)
	if err != nil {
		return nil, err
	}

	// set initial amount
	err = a.VestedAmount.Set(ctx, vestedCoin.Amount)
	if err != nil {
		return nil, err
	}

	// set denom
	err = a.VestedCoinDenom.Set(ctx, vestedCoin.Denom)
	if err != nil {
		return nil, err
	}

	// set withdrawn amount as 0
	err = a.WithdrawnAmount.Set(ctx, math.ZeroInt())
	if err != nil {
		return nil, err
	}

	return new(sdk.InitResponse), nil
}

func (a Account) WithdrawCoins(ctx *sdk.Context) (*sdk.ExecuteResponse, error) {
	beneficiary, err := a.Beneficiary.Get(ctx)
	if err != nil {
		return nil, err
	}

	if !beneficiary.Equals(ctx.Sender) {
		return nil, fmt.Errorf("unauthorized")
	}

	// check if we're after vesting start time
	startTime, err := a.StartTime.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: start time", err)
	}

	if ctx.BlockTime().Before(startTime) {
		return nil, fmt.Errorf("cannot withdraw coins before vesting start time")
	}

	withdrawable, err := a.WithrawableAmount(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: withdrawable", err)
	}
	if withdrawable.IsZero() {
		return nil, fmt.Errorf("no coins to withdraw")
	}

	// update withdrawn amount
	withdrawnSoFar, err := a.WithdrawnAmount.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: withdawn amount", err)
	}

	err = a.WithdrawnAmount.Set(ctx, withdrawnSoFar.Add(withdrawable))
	if err != nil {
		return nil, err
	}

	denom, err := a.VestedCoinDenom.Get(ctx)
	if err != nil {
		return nil, err
	}

	withdrawableCoins := sdk.NewCoin(denom, withdrawable)

	ctx.WithEvent("transferred_coins", sdk.Attribute{Key: "coins", Value: withdrawableCoins.String()})

	return new(sdk.ExecuteResponse).
		WithCosmoSDKMsg(
			&banktypes.MsgSend{
				FromAddress: ctx.Self.String(),
				ToAddress:   beneficiary.String(),
				Amount:      sdk.NewCoins(withdrawableCoins),
			},
		), nil
}

func (a Account) WithrawableAmount(ctx *sdk.Context) (math.Int, error) {
	withdrawnCoins, err := a.WithdrawnAmount.Get(ctx)
	if err != nil {
		return math.Int{}, err
	}

	unlocked, err := a.UnlockedCoins(ctx)
	if err != nil {
		return math.Int{}, err
	}

	return unlocked.Sub(withdrawnCoins), nil
}

// UnlockedCoins returns how many coins are unlocked based on the current time.
func (a Account) UnlockedCoins(ctx *sdk.Context) (math.Int, error) {
	startTime, err := a.StartTime.Get(ctx)
	if err != nil {
		return math.Int{}, err
	}
	unlocksPerSecond, err := a.UnlocksPerSecond.Get(ctx)
	if err != nil {
		return math.Int{}, err
	}

	elapsedSeconds := int64(ctx.BlockTime().Sub(startTime).Seconds())

	initialAmount, err := a.VestedAmount.Get(ctx)
	if err != nil {
		return math.Int{}, err
	}

	return math.MinInt(unlocksPerSecond.QuoRaw(elapsedSeconds), initialAmount), nil
}

func (a Account) RegisterExecuteHandler(r *sdk.ExecuteRouter) {
	sdk.RegisterExecuteHandler(r, func(ctx *sdk.Context, msg v1.MsgWithdrawUnlockedCoins) (*sdk.ExecuteResponse, error) {
		return a.WithdrawCoins(ctx)
	})
}

func (a Account) RegisterQueryHandler(r *sdk.QueryRouter) {
	sdk.RegisterQueryHandler(r, func(ctx *sdk.Context, msg v1.QueryVestingStatusRequest) (v1.QueryVestingStatusResponse, error) {
		unlockedCoins, err := a.UnlockedCoins(ctx)
		if err != nil {
			return v1.QueryVestingStatusResponse{}, err
		}

		withdrawable, err := a.WithrawableAmount(ctx)
		if err != nil {
			return v1.QueryVestingStatusResponse{}, err
		}
		return v1.QueryVestingStatusResponse{
			UnlockedAmount:     unlockedCoins,
			WithdrawableAmount: withdrawable,
		}, nil
	})
}
