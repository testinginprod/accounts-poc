package keeper

import (
	sdk2 "accounts/sdk"
	"cosmossdk.io/collections"
	"fmt"
	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"accounts/x/accounts/types"
)

type (
	Keeper struct {
		accounts map[string]sdk2.InternalAccount
		schemas  map[string]*collections.Schema

		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
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

	panic("todo")

	return &Keeper{
		accounts:   nil,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Deploy(ctx sdk.Context, kind string, deployer sdk.AccAddress, deployMsg proto.Message) (sdk.AccAddress, error) {
	account, exists := k.accounts[kind]
	if !exists {
		return nil, fmt.Errorf("unrecognized account kind: %s", kind)
	}
	panic(account)
}
