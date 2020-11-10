package qoin

var (
	isProduction    bool
	privateKey      string
	secretKey       string
	referenceNumber string
)

// SetEnvironment set the API environment
func SetEnvironment(environment string) {
	switch environment {
	case "production":
		isProduction = true
	default:
		isProduction = false
	}
}

// SetPrivateKey set the private key for generate signature
func SetPrivateKey(privKey string) {
	privateKey = privKey
}

// SetSecretKey set the secret key for generate signature
func SetSecretKey(secKey string) {
	secretKey = secKey
}

// SetReferenceNumber set the reference number for generate signature
func SetReferenceNumber(refNumber string) {
	referenceNumber = refNumber
}
