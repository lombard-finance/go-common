package address

import (
	"encoding/hex"
	"strings"

	"github.com/mr-tron/base58"
)

const (
	EVMAddressLength    = 20 // 20 bytes for EVM addresses
	SUIAddressLength    = 32 // 32 bytes for SUI addresses
	SolanaAddressLength = 32 // 32 bytes for Solana addresses
)

// IsValidBlockchainAddress verifies whether a string can represent a valid
// EVM, SUI or Solana address.
func IsValidBlockchainAddress(s string) bool {
	// Check if it could be a Solana address (no 0x prefix)
	if !strings.HasPrefix(strings.ToLower(s), "0x") {
		if res, err := base58.Decode(s); err == nil {
			return len(res) == SolanaAddressLength
		}
	}

	// Handle as potential EVM or SUI address
	s = strings.TrimPrefix(strings.TrimPrefix(s, "0x"), "0X")

	if res, err := hex.DecodeString(s); err == nil {
		return len(res) == EVMAddressLength || len(res) == SUIAddressLength
	}

	return false
}
