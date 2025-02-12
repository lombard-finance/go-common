package signature

import (
	"encoding/base64"
	"fmt"

	suimodels "github.com/block-vision/sui-go-sdk/models"
)

func SuiVerifyPersonalMessage(message string, b64Signature string) (string, bool, error) {
	// Validate base64 length
	if len(b64Signature) != SuiSignatureLengthBase64 {
		return "", false, fmt.Errorf("invalid signature length: got %d, want %d",
			len(b64Signature), SuiSignatureLengthBase64)
	}

	// Decode base64 to check the scheme flag
	sigBytes, err := base64.StdEncoding.DecodeString(b64Signature)
	if err != nil {
		return "", false, fmt.Errorf("failed to decode base64 signature: %w", err)
	}

	// Validate signature scheme
	switch sigBytes[0] {
	case SuiSigFlagEd25519:
		// Ed25519: 1 byte flag + 64 bytes signature + 32 bytes public key
		if len(sigBytes) != 97 {
			return "", false, fmt.Errorf("invalid Ed25519 signature byte length: got %d, want 97", len(sigBytes))
		}
	case SuiSigFlagSecp256k1, SuiSigFlagSecp256r1:
		// Secp256k1/r1: 1 byte flag + 64 bytes signature + 33 bytes public key
		if len(sigBytes) != 98 {
			return "", false, fmt.Errorf("invalid ECDSA signature byte length: got %d, want 98", len(sigBytes))
		}
	default:
		return "", false, fmt.Errorf("unsupported signature scheme: 0x%02x", sigBytes[0])
	}

	signer, ok, err := suimodels.VerifyPersonalMessage(message, b64Signature)
	if err != nil {
		return "", false, fmt.Errorf("verification failed: %w", err)
	}

	return signer, ok, nil
}
