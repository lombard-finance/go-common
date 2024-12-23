package signature

import (
	"encoding/hex"
	"strings"
)

func DecodeHex(str string) ([]byte, error) {
	if strings.HasPrefix(str, "0x") {
		return hex.DecodeString(strings.TrimPrefix(str, "0x"))
	}
	return hex.DecodeString(str)
}
