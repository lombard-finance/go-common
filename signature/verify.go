package signature

import (
	"bytes"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"math/big"
	"strconv"
)

func VerifySignature(signer, sign, message string) error {
	signerBytes, err := DecodeHex(signer)
	if err != nil {
		return errors.Wrap(err, "decode signer")
	}

	signature, err := DecodeHex(sign)
	if err != nil {
		return errors.Wrap(err, "decode signature")
	}

	address, err := verifyMessage(message, signature)
	if err != nil {
		return errors.Wrap(err, "verify message")
	}

	if len(address) != len(signerBytes) {
		return errors.Errorf("invalid signer length. expected %d. actual %d.", len(address), len(signerBytes))
	}

	if !bytes.Equal(address, signerBytes) {
		return errors.Errorf("wrong signer (%x != %x)", address, signerBytes)
	}

	return nil
}

func verifyMessage(message string, signedMessage []byte) ([]byte, error) {
	if len(signedMessage) != 65 {
		return nil, errors.Errorf("wrong signature length (%d != 65)", len(signedMessage))
	}

	// Hash the unsigned message using EIP-191
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)

	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if signedMessage[64] == 27 || signedMessage[64] == 28 {
		signedMessage[64] -= 27
	}

	// Extract r, s, v values
	r := new(big.Int).SetBytes(signedMessage[:32])
	s := new(big.Int).SetBytes(signedMessage[32:64])
	v := signedMessage[64]

	// Validate the signature values including `s` to be in the lower
	// half of the elliptic curve order to avoid signature malleability
	if !crypto.ValidateSignatureValues(v, r, s, true) {
		return nil, errors.New("invalid signature values")
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signedMessage)
	if err != nil {
		return nil, errors.Wrap(err, "could not recover public key from the signature")
	}
	if sigPublicKeyECDSA == nil {
		return nil, errors.New("could not get a public key from the message signature")
	}

	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).Bytes(), nil
}
