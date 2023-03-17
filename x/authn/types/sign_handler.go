package types

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SignatureHandler interface {
	VerifySignature(ctx sdk.Context, tx sdk.Tx, index int)
}

var _ SignatureHandler = (*secp256k1Credentials)(nil)

type secp256k1Credentials struct {
	PubKey        *secp256k1.PubKey
	Sequence      uint64
	AccountNumber uint64
}

func (s secp256k1Credentials) VerifySignature(ctx sdk.Context, tx sdk.Tx, index int) {
	//TODO implement me
	panic("implement me")
}
