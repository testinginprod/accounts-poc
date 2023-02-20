package sdk

import "github.com/gogo/protobuf/proto"

type Encodable[T any] interface {
	*T
	proto.Message
}
