package keeper

import (
	"accounts/sdk"
	"bytes"
	"cosmossdk.io/collections"
	"github.com/cosmos/gogoproto/proto"
	"github.com/gogo/protobuf/jsonpb"
)

type AccountsMap map[string]*InternalAccount

func (a AccountsMap) Schemas() map[string]*Schema {
	schemas := make(map[string]*Schema, len(a))
	for k, v := range a {
		schemas[k] = v.schema()
	}
	return schemas
}

func newInitMsgSchema[IM any, PIM sdk.Encodable[IM]]() *initMsgSchema {
	return &initMsgSchema{
		EncodeFromJSONToProto: func(jsonBytes []byte) (proto.Message, error) {
			x := PIM(new(IM))
			err := jsonpb.Unmarshal(bytes.NewReader(jsonBytes), x)
			return x, err
		},
		EncodeFromJSONStringToProto: func(jsonString string) (proto.Message, error) {
			x := PIM(new(IM))
			err := jsonpb.UnmarshalString(jsonString, x)
			return x, err
		},
	}
}

type initMsgSchema struct {
	EncodeFromJSONToProto       func(jsonBytes []byte) (proto.Message, error)
	EncodeFromJSONStringToProto func(jsonString string) (proto.Message, error)
}

type Schema struct {
	InitMsg  *initMsgSchema
	messages interface{}
	queries  interface{}
	state    *collections.Schema
}

type runtime struct {
	accounts map[string]InternalAccount
	schemas  map[string]*Schema
}
