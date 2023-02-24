package sdk

import (
	"cosmossdk.io/collections/codec"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"time"
)

type Encodable[T any] interface {
	*T
	proto.Message
}

var IdentityFromString = sdk.AccAddressFromBech32
var IdentityKey = codec.NewBytesKey[Identity]()
var IdentityValue = codec.KeyToValueCodec(IdentityKey)
var IntKey codec.KeyCodec[math.Int] = nil
var IntValue = codec.KeyToValueCodec(IntKey)
var TimeValue codec.ValueCodec[time.Time] = nil
