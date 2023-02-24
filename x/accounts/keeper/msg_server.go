package keeper

import (
	"accounts/utils"
	"accounts/x/accounts/types"
	"context"
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	k Keeper
}

func (m msgServer) Deploy(ctx context.Context, deploy *types.MsgDeploy) (*types.MsgDeployResponse, error) {
	addr, err := sdk.AccAddressFromBech32(deploy.Sender)
	if err != nil {
		return nil, err
	}

	initMsg, err := utils.UnmarshalAnyBytes(deploy.InitMessage)
	if err != nil {
		return nil, fmt.Errorf("unmarshal any error: %w", err)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	accountAddr, accountID, data, err := m.k.Deploy(sdkCtx, deploy.Kind, addr, deploy.Funds, initMsg)
	if err != nil {
		return nil, err
	}

	var dataAny *codectypes.Any

	if data != nil {
		dataAny, err = codectypes.NewAnyWithValue(data)
		if err != nil {
			return nil, err
		}
	}

	sdkCtx.EventManager().EmitEvent(sdk.NewEvent("account_deployed", sdk.NewAttribute("address", accountAddr.String())))

	return &types.MsgDeployResponse{
		Address: accountAddr.String(),
		Id:      accountID,
		Data:    dataAny,
	}, nil
}

func (m msgServer) Execute(ctx context.Context, execute *types.MsgExecute) (*types.MsgExecuteResponse, error) {
	addr, err := sdk.AccAddressFromBech32(execute.Address)
	if err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(execute.Sender)
	if err != nil {
		return nil, err
	}

	executeMsg, err := utils.UnmarshalAnyBytes(execute.Message)
	if err != nil {
		return nil, err
	}

	data, err := m.k.Execute(sdk.UnwrapSDKContext(ctx), sender, addr, execute.Funds, executeMsg)
	if err != nil {
		return nil, err
	}

	var anyData *codectypes.Any
	if data != nil {
		anyData, err = codectypes.NewAnyWithValue(data)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgExecuteResponse{Data: anyData}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

var _ types.MsgServer = msgServer{}
