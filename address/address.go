package address

import (
	"fmt"
)

const (
	EVMAddressLength = 40 // 40 hex chars for EVM addresses
	SUIAddressLength = 64 // 64 hex chars for SUI addresses
)

// IsValidBlockchainAddress verifies whether a string can represent a valid hex-encoded
// EVM or SUI address.
func IsValidBlockchainAddress(s string) (bool, string) {
	// Remove 0x prefix if present
	if has0xPrefix(s) {
		s = s[2:]
	}

	// Check if it's hex
	if !isHex(s) {
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

// has0xPrefix validates str begins with '0x' or '0X'.
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}
