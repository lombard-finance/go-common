package signature

import (
	"encoding/json"
	"errors"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/holiman/uint256"
	"github.com/storyicon/sigverify"
)

var (
	ErrZeroMaxMintFee = errors.New("mint fee cannot be 0")
	ErrMissingFee     = errors.New("fee field is required")
)

type Domain struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	ChainId           int    `json:"chainId"`
	VerifyingContract string `json:"verifyingContract"`
}

type Message struct {
	Expiry  int64        `json:"expiry"`
	ChainId uint64       `json:"chainId"`
	Fee     *uint256.Int `json:"fee"`
}

type Type struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Types map[string][]Type

type TypedData struct {
	Types       Types   `json:"types"`
	PrimaryType string  `json:"primaryType"`
	Domain      Domain  `json:"domain"`
	Message     Message `json:"message"`
}

func ExtractTypedDataValues(typedDataStr string) (time.Time, uint64, uint64, error) {
	var data TypedData
	if err := json.Unmarshal([]byte(typedDataStr), &data); err != nil {
		return time.Time{}, 0, 0, err
	}

	if data.Message.Fee == nil {
		return time.Time{}, 0, 0, ErrMissingFee
	}

	if data.Message.Fee.Cmp(uint256.NewInt(0)) == 0 {
		return time.Time{}, 0, 0, ErrZeroMaxMintFee
	}
	return time.Unix(data.Message.Expiry, 0), data.Message.Fee.Uint64(), data.Message.ChainId, nil
}

func VerifyEIP712Signature(signerHex, signatureHex, typedDataMarshaled string) error {
	var data apitypes.TypedData
	if err := json.Unmarshal([]byte(typedDataMarshaled), &data); err != nil {
		return err
	}

	signerEqualAddress, err := sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress(signerHex),
		data,
		signatureHex,
	)
	if err != nil {
		return err
	}

	if !signerEqualAddress {
		return errors.New("signer does not match recovered address")
	}

	return nil
}
