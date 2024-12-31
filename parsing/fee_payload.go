package parsing

import (
	"math/big"

	"github.com/pkg/errors"
)

// extract mint fee from the fee payload
func ExtractMintFeeFromFeePayload(feePayload []byte) (uint64, error) {
	// Check minimum length (4 bytes function signature + 32 bytes fee + 32 bytes expiry)
	if len(feePayload) < 68 || len(feePayload) > 68 {
		return 0, errors.Errorf("invalid fee payload length: %d", len(feePayload))
	}

	// Skip function selector (first 4 bytes) and extract fee amount
	feeAmount := new(big.Int).SetBytes(feePayload[4:36])

	return feeAmount.Uint64(), nil
}
