package sdk

import (
	"context"
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
)

func NewSchema() *collections.SchemaBuilder { return collections.NewSchemaBuilder(storeService{}) }

type storeService struct{}

func (storeService) OpenKVStore(ctx context.Context) store.KVStore { return ctx.(*Context).Store }
