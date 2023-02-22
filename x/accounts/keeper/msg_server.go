package keeper

import (
	"accounts/x/accounts/types"
	"context"
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	types2 "github.com/gogo/protobuf/types"
	"reflect"
)

type msgServer struct {
	k Keeper
}

func (m msgServer) Deploy(ctx context.Context, deploy *types.MsgDeploy) (*types.MsgDeployResponse, error) {
	addr, err := sdk.AccAddressFromBech32(deploy.Sender)
	if err != nil {
		return nil, err
	}

	initMsg, err := unmarshalAny(deploy.InitMessage)
	if err != nil {
		return nil, err
	}

	accountAddr, accountID, data, err := m.k.Deploy(sdk.UnwrapSDKContext(ctx), deploy.Kind, addr, deploy.Funds, initMsg)
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

	executeMsg, err := unmarshalAny(execute.Message)
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

func unmarshalAny(any *codectypes.Any) (proto.Message, error) {
	messageName, err := types2.AnyMessageName(&types2.Any{TypeUrl: any.TypeUrl})
	if err != nil {
		return nil, err
	}
	messageType := proto.MessageType(messageName)
	if messageType == nil {
		return nil, fmt.Errorf("unknown message: %s", messageName)
	}

	msg := reflect.New(messageType.Elem()).Interface().(proto.Message)
	return msg, proto.Unmarshal(any.Value, msg)
}
