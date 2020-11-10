package qoin

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// GenerateSignature returns generated RSA signature
func GenerateSignature(text string) interface{} {
	rng := rand.Reader
	message := []byte(text)
	hashed := sha256.Sum256(message)
	key, _ := rsa.GenerateKey(rng, 1024)
	signature, _ := rsa.SignPKCS1v15(rng, key, crypto.SHA256, hashed[:])

	return signature
}
