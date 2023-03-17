package sdk

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/collections/codec"
	"cosmossdk.io/math"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"time"
)

type Encodable[T any] interface {
	*T
	proto.Message
}

var AccAddressFromBech32 = sdk.AccAddressFromBech32
var IdentityKey = codec.NewBytesKey[AccAddress]()
var AccAddressValue = codec.KeyToValueCodec(IdentityKey)
var IntValue codec.ValueCodec[math.Int] = intValueCodec{}
var TimeKey codec.KeyCodec[time.Time] = timeKey{}
var TimeValue = codec.KeyToValueCodec(TimeKey)

type intValueCodec struct{}

func (i intValueCodec) Encode(value math.Int) ([]byte, error) {
	return value.Marshal()
}

func (i intValueCodec) Decode(b []byte) (math.Int, error) {
	v := new(math.Int)
	err := v.Unmarshal(b)
	if err != nil {
		return math.Int{}, err
	}
	return *v, nil
}

func (i intValueCodec) EncodeJSON(value math.Int) ([]byte, error) {
	return value.MarshalJSON()
}

func (i intValueCodec) DecodeJSON(b []byte) (math.Int, error) {
	v := new(math.Int)
	err := v.UnmarshalJSON(b)
	if err != nil {
		return math.Int{}, err
	}
	return *v, nil
}

func (i intValueCodec) Stringify(value math.Int) string {
	return value.String()
}

func (i intValueCodec) ValueType() string {
	return "math.Int"
}

type timeKey struct{}

func (t timeKey) Encode(buffer []byte, key time.Time) (int, error) {
	return collections.Int64Key.Encode(buffer, key.UnixMilli())
}

func (t timeKey) Decode(buffer []byte) (int, time.Time, error) {
	r, i, err := collections.Int64Key.Decode(buffer)
	if err != nil {
		return 0, time.Time{}, err
	}
	return r, time.UnixMilli(i), nil
}

func (t timeKey) Size(key time.Time) int {
	return collections.Int64Key.Size(0)
}

func (t timeKey) EncodeJSON(value time.Time) ([]byte, error) {
	return json.Marshal(value)
}

func (t timeKey) DecodeJSON(b []byte) (time.Time, error) {
	v := new(time.Time)
	err := json.Unmarshal(b, v)
	if err != nil {
		return time.Time{}, err
	}
	return *v, nil
}

func (t timeKey) Stringify(key time.Time) string {
	return key.String()
}

func (t timeKey) KeyType() string {
	return "time.Time"
}

func (t timeKey) EncodeNonTerminal(buffer []byte, key time.Time) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t timeKey) DecodeNonTerminal(buffer []byte) (int, time.Time, error) {
	//TODO implement me
	panic("implement me")
}

func (t timeKey) SizeNonTerminal(key time.Time) int {
	//TODO implement me
	panic("implement me")
}
