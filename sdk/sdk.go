package sdk

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type Identity = sdk.AccAddress

type Account[T any, PT Encodable[T]] interface {
	Init(ctx *Context, t T) (*InitResponse, error)
	RegisterExecuteHandler(router *ExecuteRouter)
	RegisterQueryHandler(router *QueryRouter)
}

func RegisterExecuteHandler[T any, PT Encodable[T]](router *ExecuteRouter, handler func(ctx *Context, msg T) (*ExecuteResponse, error)) {
	name := proto.MessageName(PT(new(T)))
	h := func(ctx *Context, m proto.Message) (*ExecuteResponse, error) {
		concrete, ok := m.(PT)
		if ok {
			return nil, fmt.Errorf("routing error")
		}
		return handler(ctx, *concrete)
	}

	router.handlers[name] = h
}

func NewExecuteRouter() *ExecuteRouter {
	return &ExecuteRouter{handlers: map[string]func(ctx *Context, msg proto.Message) (*ExecuteResponse, error){}}
}

type ExecuteRouter struct {
	handlers map[string]func(ctx *Context, msg proto.Message) (*ExecuteResponse, error)
}

func (e *ExecuteRouter) Handler() func(ctx *Context, msg proto.Message) (*ExecuteResponse, error) {
	return func(ctx *Context, msg proto.Message) (*ExecuteResponse, error) {
		name := proto.MessageName(msg)
		handler, exist := e.handlers[name]
		if !exist {
			return nil, fmt.Errorf("unknown execute request: %s", name)
		}
		return handler(ctx, msg)
	}
}

func NewQueryRouter() *QueryRouter {
	return &QueryRouter{handlers: map[string]func(ctx *Context, msg proto.Message) (proto.Message, error){}}
}

type QueryRouter struct {
	handlers map[string]func(ctx *Context, msg proto.Message) (proto.Message, error)
}

func (q *QueryRouter) Handler() func(ctx *Context, msg proto.Message) (proto.Message, error) {
	return func(ctx *Context, msg proto.Message) (proto.Message, error) {
		name := proto.MessageName(msg)
		handler, exist := q.handlers[name]
		if !exist {
			return nil, fmt.Errorf("unknown query request: %s", name)
		}
		return handler(ctx, msg)
	}
}

func RegisterQueryHandler[Req, Resp any, PReq Encodable[Req], PResp Encodable[Resp]](r *QueryRouter, handler func(ctx *Context, msg Req) (Resp, error)) {
	reqName := proto.MessageName(PReq(new(Req)))
	r.handlers[reqName] = func(ctx *Context, msg proto.Message) (proto.Message, error) {
		concrete, ok := msg.(PReq)
		if !ok {
			return nil, fmt.Errorf("invalid query request: wanted %T, got %T", new(Req), msg)
		}

		resp, err := handler(ctx, *concrete)
		if err != nil {
			return nil, err
		}

		return PResp(&resp), nil
	}
}
