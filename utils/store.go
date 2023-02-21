package utils

import (
	"context"
	"cosmossdk.io/core/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type KVStoreOpenerFunc func(ctx context.Context) store.KVStore

func (f KVStoreOpenerFunc) OpenKVStore(ctx context.Context) store.KVStore { return f(ctx) }

func SDKStoreToCoreStoreService(key storetypes.StoreKey) store.KVStoreService {
	return coreStoreService{key: key}
}

type coreStoreService struct {
	key storetypes.StoreKey
}

func (c coreStoreService) OpenKVStore(ctx context.Context) store.KVStore {
	return SDKStoreToCoreStore(sdk.UnwrapSDKContext(ctx).KVStore(c.key))
}

func SDKStoreToCoreStore(store storetypes.KVStore) store.KVStore {
	return coreStore{sdkStore: store}
}

type coreStore struct {
	sdkStore storetypes.KVStore
}

func (c coreStore) Get(key []byte) ([]byte, error) {
	return c.sdkStore.Get(key), nil
}

func (c coreStore) Has(key []byte) (bool, error) {
	return c.sdkStore.Has(key), nil
}

func (c coreStore) Set(key, value []byte) error {
	c.sdkStore.Set(key, value)
	return nil
}

func (c coreStore) Delete(key []byte) error {
	c.sdkStore.Delete(key)
	return nil
}

func (c coreStore) Iterator(start, end []byte) (store.Iterator, error) {
	return c.sdkStore.Iterator(start, end), nil
}

func (c coreStore) ReverseIterator(start, end []byte) (store.Iterator, error) {
	return c.sdkStore.ReverseIterator(start, end), nil
}
