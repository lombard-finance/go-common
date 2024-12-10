package common

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func HexToBase64(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(strings.TrimPrefix(hexStr, "0x"))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func StringToBase64(str string) string {
	bytes := []byte(str)
	return base64.StdEncoding.EncodeToString(bytes)
}

func Base64ToHex(base64Str string) (string, error) {
	addrbytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(addrbytes), nil
}
