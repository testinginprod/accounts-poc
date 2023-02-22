package sdk

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

type accountMsg struct {
	dst Identity
	msg proto.Message
}

type message struct {
	moduleMsg  sdk.Msg
	accountMsg *accountMsg
}

type InitResponse struct{}

type ExecuteResponse struct {
	messages []*message
}

func (r *ExecuteResponse) WithCosmoSDKMsg(msg sdk.Msg) *ExecuteResponse {
	r.messages = append(r.messages, &message{
		moduleMsg: msg,
	})

	return r
}
