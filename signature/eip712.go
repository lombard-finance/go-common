package signature

import (
	"encoding/json"
	"errors"
	"time"
)

var ErrZeroMaxMintFee = errors.New("mint fee cannot be 0")

type Message struct {
	Expiry  int64  `json:"expiry"`
	ChainId uint64 `json:"chainId"`
	Fee     uint64 `json:"fee"`
}

type TypedData struct {
	Message `json:"message"`
}

func ExtractTypedDataValues(typedDataStr string) (expiryTime time.Time, feeApproved uint64, chainId uint64, err error) {
	var data TypedData
	if err := json.Unmarshal([]byte(typedDataStr), &data); err != nil {
		return time.Time{}, 0, 0, err
	}

	if data.Message.Fee == 0 {
		return time.Time{}, 0, 0, ErrZeroMaxMintFee
	}

	expiryTime = time.Unix(data.Message.Expiry, 0)
	chainId = data.Message.ChainId
	return expiryTime, feeApproved, chainId, nil
}

func VerifyEIP712Signature(signerBytes, signature, typedData string) error {
	return VerifySignature(signerBytes, signature, typedData)
}
