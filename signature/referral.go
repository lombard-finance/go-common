package signature

const ReferralMessageTemplate = "I have read and agreed to the terms of service: https://docs.lombard.finance/legals/terms-of-service"

func VerifyReferralSignature(signer, signature []byte) error {
	return VerifySignature(signer, signature, ReferralMessageTemplate)
}
