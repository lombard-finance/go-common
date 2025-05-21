package address

import (
	"encoding/hex"
	"strings"

	"regexp"

	"github.com/mr-tron/base58"
)

const (
	EVMAddressLength    = 20 // 20 bytes for EVM addresses
	SUIAddressLength    = 32 // 32 bytes for SUI addresses
	SolanaAddressLength = 32 // 32 bytes for Solana addresses
)

// hexPattern matches valid hexadecimal strings with optional 0x prefix
var hexPattern = regexp.MustCompile(`^(?:0x)?[0-9a-fA-F]+$`)

// NormalizeAddress normalizes addresses by lowercasing hex addresses
// and leaving other formats (like Base58) unchanged.
func NormalizeAddress(address string) string {
	if hexPattern.MatchString(address) {
		return strings.ToLower(address)
	}
	return address
}

// IsValidDestinationBlockchainAddress verifies whether a string can represent a valid
// EVM, SUI or Solana address.
func IsValidDestinationBlockchainAddress(s string) bool {
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
