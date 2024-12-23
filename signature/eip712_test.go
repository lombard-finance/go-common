package signature

import (
	"encoding/json"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyEIPSignature(t *testing.T) {
	const (
		validSign   = "0xbb3bb350040ea819cfcf10604cfd7e5eef46b49b0700c70131d69f08f958ed255e70c72c7c0ed41edd62f897df0cc32678da9d745e49f73a5daa0d1f0bfda80d1b"
		invalidSign = "0xbb3bb350040ea819cfcf10604cfd7e5eef46b49b0700"
		signer      = "0x8C6bF4b04363910443cCc8F3B71B267EC3b96241"
	)

	typeData := TypedData{
		Domain: Domain{
			Name:              "Lombard Staked Bitcoin",
			Version:           "1",
			ChainId:           11155111,
			VerifyingContract: "0xc47e4b3124597FDF8DD07843D4a7052F2eE80C30",
		},
		Message: Message{
			Expiry: 1735043332, ChainId: 11155111, Fee: uint256.NewInt(1)},
		PrimaryType: "feeApproval",
		Types: map[string][]Type{
			"EIP712Domain": {{Name: "name", Type: "string"}, {Name: "version", Type: "string"}, {Name: "chainId", Type: "uint256"}, {Name: "verifyingContract", Type: "address"}},
			"feeApproval":  {{Name: "chainId", Type: "uint256"}, {Name: "fee", Type: "uint256"}, {Name: "expiry", Type: "uint256"}},
		},
	}

	validMsg, err := json.Marshal(&typeData)
	assert.NoError(t, err)

	t.Run("valid signature", func(t *testing.T) {
		assert.NoError(t, VerifyEIP712Signature(signer, validSign, string(validMsg)))
	})

	t.Run("invalid message", func(t *testing.T) {
		typeData := TypedData{
			Message: Message{
				Expiry: 0, ChainId: 11155111, Fee: uint256.NewInt(1)},
		}
		data, err := json.Marshal(&typeData)
		assert.NoError(t, err)

		assert.Error(t, VerifyEIP712Signature(signer, validSign, string(data)))
	})

	t.Run("invalid signature", func(t *testing.T) {
		assert.Error(t, VerifyEIP712Signature(signer, invalidSign, string(validMsg)))
	})
}
