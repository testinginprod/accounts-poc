package keeper

import (
	sdk2 "accounts/sdk"
	"accounts/utils"
	"context"
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"accounts/x/accounts/types"
)

var AccountIDPrefix = collections.NewPrefix([]byte{0x0, 0x0})
var accountsStoragePrefix = []byte{0x0} // for versioning purposes

type (
	Keeper struct {
		accounts map[string]*InternalAccount

		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		AccountID collections.Sequence
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accounts Accounts,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	moduleSchema := collections.NewSchemaBuilder(utils.SDKStoreToCoreStoreService(storeKey))
	accMap := map[string]*InternalAccount{}
	for _, accCreator := range accounts {
		accountSchema := collections.NewSchemaBuilder(utils.KVStoreOpenerFunc(func(ctx context.Context) store.KVStore {
			return ctx.(*sdk2.Context).Store
		}))

		account, err := accCreator(accountSchema)
		if err != nil {
			panic(err)
		}
		name := account.name()
		if _, exists := accMap[name]; exists {
			panic("already registered account " + name)
		}
		accMap[name] = account
	}

	return &Keeper{
		accounts:   accMap,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		AccountID:  collections.NewSequence(moduleSchema, AccountIDPrefix, "account_id"),
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Deploy(ctx sdk.Context, kind string, deployer sdk.AccAddress, deployMsg proto.Message) (sdk.AccAddress, uint64, proto.Message, error) {
	account, exists := k.accounts[kind]
	if !exists {
		return nil, 0, nil, fmt.Errorf("unknown account kind")
	}

	accountID, err := k.AccountID.Next(ctx)
	if err != nil {
		return nil, 0, nil, err
	}

	accountAddr := authtypes.NewModuleAddress(fmt.Sprintf("%s/%d", "accounts", accountID))

	accCtx := sdk2.NewContextFromSDK(ctx, deployer, accountAddr, accountID, k.accountStore(ctx, accountID))

	_, err = account.init(accCtx, deployMsg) // TODO routing of resp
	if err != nil {
		return nil, 0, nil, err
	}

	return accountAddr, accountID, nil /*TODO*/, nil

}

func (k Keeper) accountStore(ctx sdk.Context, id uint64) store.KVStore {
	return utils.SDKStoreToCoreStore(
		prefix.NewStore(
			ctx.KVStore(k.storeKey),
			append(accountsStoragePrefix, sdk.Uint64ToBigEndian(id)...),
		),
	)
}
