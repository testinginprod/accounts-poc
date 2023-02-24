package sdk

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type AccountMsg struct {
	Dst   Identity
	Msg   proto.Message
	Funds Coins
}

// Message acts as a oneof. There can be only one of
type Message struct {
	module  sdk.Msg
	account *AccountMsg
}

func (m *Message) IsModuleMsg() bool {
	return m.module != nil
}

func (m *Message) IsAccountMsg() bool {
	return m.account != nil
}

func (m *Message) AccountMsg() *AccountMsg {
	return m.account
}

func (m *Message) ModuleMsg() sdk.Msg {
	return m.module
}

type InitResponse struct {
}

type ExecuteResponse struct {
	messages []*Message
}

func (r *ExecuteResponse) WithCosmoSDKMsg(msg sdk.Msg) *ExecuteResponse {
	r.messages = append(r.messages, &Message{
		module: msg,
	})
	return r
}

func (r *ExecuteResponse) WithAccountMsg(dst Identity, msg proto.Message) *ExecuteResponse {
	r.messages = append(r.messages, &Message{account: &AccountMsg{
		Dst: dst,
		Msg: msg,
	}})

	return r
}

func (r *ExecuteResponse) Messages() []*Message {
	return r.messages
}
