package types

const (
	// ModuleName defines the module name
	ModuleName = "accounts"

	// StoreKey defines the primary module store key
	StoreKey = "something"

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_accounts"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
