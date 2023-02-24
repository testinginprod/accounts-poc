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
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"accounts/x/accounts/types"
)

var (
	AccountIDPrefix      = collections.NewPrefix([]byte{0x0, 0x0})
	AccountsByKindPrefix = collections.NewPrefix([]byte{0x0, 0x1})
	AccountsByIDPrefix   = collections.NewPrefix([]byte{0x0, 0x2})

	accountsStoragePrefix = []byte{0x1} // for versioning purposes
)

type MsgRouter interface {
	Handler(msg sdk.Msg) func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error)
}

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	memKey     storetypes.StoreKey
	paramstore paramtypes.Subspace

	accounts map[string]*InternalAccount
	router   MsgRouter

	AccountID      collections.Sequence
	AccountsByKind collections.Map[sdk2.Identity, string] // MAYBE use an indexed map.
	AccountsByID   collections.Map[sdk2.Identity, uint64]
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	router MsgRouter,
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
		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		paramstore:     ps,
		accounts:       accMap,
		router:         router,
		AccountID:      collections.NewSequence(moduleSchema, AccountIDPrefix, "account_id"),
		AccountsByKind: collections.NewMap(moduleSchema, AccountsByKindPrefix, "accounts_by_kind", sdk2.IdentityKey, collections.StringValue),
		AccountsByID:   collections.NewMap(moduleSchema, AccountsByIDPrefix, "accounts_by_id", sdk2.IdentityKey, collections.Uint64Value),
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Deploy(ctx sdk.Context, kind string, deployer sdk.AccAddress, funds sdk.Coins, deployMsg proto.Message) (sdk.AccAddress, uint64, proto.Message, error) {
	account, exists := k.accounts[kind]
	if !exists {
		return nil, 0, nil, fmt.Errorf("unknown account kind")
	}

	accountID, err := k.AccountID.Next(ctx)
	if err != nil {
		return nil, 0, nil, err
	}

	accountAddr := authtypes.NewModuleAddress(fmt.Sprintf("%s/%d", "accounts", accountID))

	if err = k.depositFunds(ctx, deployer, accountAddr, funds); err != nil {
		return nil, 0, nil, err
	}

	accCtx := sdk2.NewContextFromSDK(ctx, deployer, accountAddr, accountID, k.accountStore(ctx, accountID), funds)

	_, err = account.init(accCtx, deployMsg) // TODO routing of resp
	if err != nil {
		return nil, 0, nil, err
	}

	if err = k.AccountsByKind.Set(ctx, accountAddr, kind); err != nil {
		return nil, 0, nil, err
	}
	if err = k.AccountsByID.Set(ctx, accountAddr, accountID); err != nil {
		return nil, 0, nil, err
	}

	return accountAddr, accountID, nil /*TODO*/, nil
}

func (k Keeper) Execute(ctx sdk.Context, from sdk2.Identity, to sdk2.Identity, funds sdk2.Coins, msg proto.Message) (proto.Message, error) {
	kind, err := k.AccountsByKind.Get(ctx, to)
	if err != nil {
		return nil, fmt.Errorf("unknown account identifier: %s", to)
	}

	account, ok := k.accounts[kind]
	if !ok {
		return nil, fmt.Errorf("unknown account kind: %s", kind)
	}

	id, err := k.AccountsByID.Get(ctx, from)
	if err != nil {
		return nil, err
	}

	if err = k.depositFunds(ctx, from, to, funds); err != nil {
		return nil, err
	}

	accCtx := sdk2.NewContextFromSDK(ctx, from, to, id, k.accountStore(ctx, id), funds)

	resp, err := account.execute(accCtx, msg)
	if err != nil {
		return nil, err
	}

	err = k.route(ctx, to, resp.Messages())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (k Keeper) route(ctx sdk.Context, from sdk2.Identity, resp []*sdk2.Message) error {
	for _, msg := range resp {
		switch {
		case msg.IsAccountMsg():
			accountMsg := msg.AccountMsg()
			_, err := k.Execute(ctx, from, accountMsg.Dst, accountMsg.Funds, accountMsg.Msg)
			if err != nil {
				return err
			}
		case msg.IsModuleMsg():
			moduleMsg := msg.ModuleMsg()
			signers := moduleMsg.GetSigners()
			if len(signers) != 1 {
				return fmt.Errorf("invalid signers: wanted only 1")
			}
			if !signers[0].Equals(from) {
				return fmt.Errorf("invalid signers: message can only be signed by %s, got: %s", from, signers[0].String())
			}
			handler := k.router.Handler(moduleMsg)
			if handler == nil {
				return fmt.Errorf("unknown module msg: %T", msg)
			}
			err := moduleMsg.ValidateBasic()
			if err != nil {
				return fmt.Errorf("invalid module message: %w", err)
			}
			_, err = handler(ctx, moduleMsg)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("empty message")
		}
	}
	return nil
}

func (k Keeper) depositFunds(ctx sdk.Context, from sdk2.Identity, to sdk2.Identity, amount sdk2.Coins) error {
	if amount.IsZero() {
		return nil
	}
	// TODO: we can create a fast path here by saving the handler for the bank message directly into the keeper.
	msg := &banktypes.MsgSend{
		FromAddress: from.String(),
		ToAddress:   to.String(),
		Amount:      amount,
	}
	_, err := k.router.Handler(msg)(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) accountStore(ctx sdk.Context, id uint64) store.KVStore {
	return utils.SDKStoreToCoreStore(
		prefix.NewStore(
			ctx.KVStore(k.storeKey),
			append(accountsStoragePrefix, sdk.Uint64ToBigEndian(id)...),
		),
	)
}
