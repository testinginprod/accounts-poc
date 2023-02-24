package sdk

import (
	"bytes"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

func NewMsgSchema[IM any, PIM Encodable[IM]]() *MsgSchema {
	return &MsgSchema{
		UnmarshalJSONBytes: func(jsonBytes []byte) (proto.Message, error) {
			x := PIM(new(IM))
			err := jsonpb.Unmarshal(bytes.NewReader(jsonBytes), x)
			return x, err
		},
		UnmarshalJSONString: func(jsonString string) (proto.Message, error) {
			x := PIM(new(IM))
			err := jsonpb.UnmarshalString(jsonString, x)
			return x, err
		},
	}
}

type MsgSchema struct {
	UnmarshalJSONBytes  func(bytes []byte) (proto.Message, error)
	UnmarshalJSONString func(str string) (proto.Message, error)
}
