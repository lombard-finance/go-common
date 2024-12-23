package signature

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyEIPSignature(t *testing.T) {
	const validSign = "0xbb3bb350040ea819cfcf10604cfd7e5eef46b49b0700c70131d69f08f958ed255e70c72c7c0ed41edd62f897df0cc32678da9d745e49f73a5daa0d1f0bfda80d1b"
	t.Run("valid signature", func(t *testing.T) {
		typeData := TypedData{
			Message: Message{
				Expiry: 1735043332, ChainId: 11155111, Fee: 1},
		}
		data, err := json.Marshal(&typeData)
		assert.NoError(t, err)

		assert.NoError(t, VerifyEIP712Signature("signer", validSign, string(data))
	})
}
