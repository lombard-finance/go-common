package signature

import (
	"github.com/gagliardetto/solana-go"
	"github.com/pkg/errors"
)

func VerifySolanaSignature(message, base58Signature, base58Address string) (bool, error) {
	signature, err := solana.SignatureFromBase58(base58Signature)
	if err != nil {
		return false, errors.Wrap(err, "failed to get signature")
	}

	publicKey, err := solana.PublicKeyFromBase58(base58Address)
	if err != nil {
		return false, errors.Wrap(err, "failed to get public key")
	}

	return publicKey.Verify([]byte(message), signature), nil
}
