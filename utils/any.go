package utils

import (
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	prototypes "github.com/gogo/protobuf/types"
	"reflect"
)

func UnmarshalAnyBytes(b []byte) (proto.Message, error) {
	any := new(codectypes.Any)
	err := proto.Unmarshal(b, any)
	if err != nil {
		return nil, err
	}
	messageName, err := prototypes.AnyMessageName(&prototypes.Any{TypeUrl: any.TypeUrl})
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

func MarshalAnyBytes(m proto.Message) ([]byte, error) {
	any, err := codectypes.NewAnyWithValue(m)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(any)
}
