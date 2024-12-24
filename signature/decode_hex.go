package signature

import (
	"encoding/hex"
	"strings"
)

func DecodeHex(str string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(str, "0x"))
}
