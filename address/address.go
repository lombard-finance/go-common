package address

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	EVMAddressLength = 40 // 40 hex chars for EVM addresses
	SUIAddressLength = 64 // 64 hex chars for SUI addresses
)

// IsValidBlockchainAddress verifies whether a string can represent a valid hex-encoded
// EVM or SUI address.
func IsValidBlockchainAddress(s string) (bool, string) {
	s = strings.TrimPrefix(strings.TrimPrefix(s, "0x"), "0X")

	if _, err := hex.DecodeString(s); err != nil {
		return false, "not a hex string"
	}

	// Check for valid lengths
	switch len(s) {
	case EVMAddressLength:
		return true, "evm"
	case SUIAddressLength:
		return true, "sui"
	default:
		return false, fmt.Sprintf("invalid address length: got %d hex chars, expected %d (EVM) or %d (SUI)",
			len(s), EVMAddressLength, SUIAddressLength)
	}
}
