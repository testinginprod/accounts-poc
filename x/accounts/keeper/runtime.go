package keeper

import (
	"accounts/sdk"
	"bytes"
	"cosmossdk.io/collections"
	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/cosmos/gogoproto/proto"
)

type AccountsMap map[string]*InternalAccount

func (a AccountsMap) Schemas() map[string]*Schema {
	schemas := make(map[string]*Schema, len(a))
	for k, v := range a {
		schemas[k] = v.schema()
	}
	return schemas
}

func newInitMsgSchema[IM any, PIM sdk.Encodable[IM]]() *msgSchema {
	return &msgSchema{
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

type msgSchema struct {
	UnmarshalJSONBytes  func(jsonBytes []byte) (proto.Message, error)
	UnmarshalJSONString func(jsonString string) (proto.Message, error)
}

type Schema struct {
	InitMsg     *msgSchema
	ExecuteMsgs map[string]*sdk.MsgSchema
	queries     map[string]interface{}
	state       *collections.Schema
}

type runtime struct {
	accounts map[string]InternalAccount
	schemas  map[string]*Schema
}
