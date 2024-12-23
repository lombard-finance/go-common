package signature

import (
	"encoding/json"
	"errors"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/holiman/uint256"
	"github.com/storyicon/sigverify"
	"time"
)

var ErrZeroMaxMintFee = errors.New("mint fee cannot be 0")

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

func ExtractTypedDataValues(typedDataStr string) (expiryTime time.Time, feeApproved uint64, chainId uint64, err error) {
	var data TypedData
	if err := json.Unmarshal([]byte(typedDataStr), &data); err != nil {
		return time.Time{}, 0, 0, err
	}

	if data.Message.Fee.Cmp(uint256.NewInt(0)) == 0 {
		return time.Time{}, 0, 0, ErrZeroMaxMintFee
	}

	expiryTime = time.Unix(data.Message.Expiry, 0)
	chainId = data.Message.ChainId
	return expiryTime, feeApproved, chainId, nil
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
