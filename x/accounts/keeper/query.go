package keeper

import (
	"accounts/x/accounts/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) AccountKind(ctx context.Context, request *types.QueryAccountKindRequest) (*types.QueryAccountKindResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, err
	}

	kind, err := k.AccountsKind.Get(ctx, addr)
	if err != nil {
		return nil, err
	}

	return &types.QueryAccountKindResponse{Kind: kind}, nil
}
