package utils

import (
	"encoding/hex"

	"github.com/pkg/errors"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

func ExtractPkScriptAddrs(scriptPubkey string) (string, error) {
	// Decode the hex-encoded scriptPubkey
	pkScript, err := hex.DecodeString(scriptPubkey)
	if err != nil {
		return "", errors.Wrap(err, "error decoding hex string")
	}

	// Extract the address from the pkScript
	_, addresses, _, err := txscript.ExtractPkScriptAddrs(pkScript, &chaincfg.MainNetParams)
	if err != nil {
		return "", errors.Wrap(err, "error extracting addresses")
	}

	if len(addresses) > 0 {
		return addresses[0].String(), nil
	} else {
		return "", errors.New("no addresses found")
	}
}
