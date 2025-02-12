package signature

const (
	// EvmEcdsaSignatureLength represents length of standard EVM ECDSA signature
	// (32 bytes R + 32 bytes S + 1 byte V)
	EvmEcdsaSignatureLength = 65
	// EvmSignatureLengthHex represents length of hex-encoded standard EVM ECDSA signature
	// EvmEcdsaSignatureLength * 2 (for hex)
	EvmEcdsaSignatureLengthHex = 130

	// Sui signature scheme flags
	SuiSigFlagEd25519   byte = 0x00
	SuiSigFlagSecp256k1 byte = 0x01
	SuiSigFlagSecp256r1 byte = 0x02

	// All valid SUI signatures in base64 are 132 characters
	SuiSignatureLengthBase64 = 132
)
