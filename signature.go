package qoin

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// GenerateSignature returns generated RSA signature
func GenerateSignature(text string) string {
	signer, err := loadPrivateKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Signer is damaged: %s", err)
	}

	signed, err := signer.Sign([]byte(text))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not sign request: %s", err)
	}

	return base64.StdEncoding.EncodeToString(signed)
}

func loadPrivateKey() (Signer, error) {
	return parsePrivateKey([]byte(privateKey))
}

func parsePrivateKey(pemBytes []byte) (Signer, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("SSH: No key found")
	}

	var rawKey interface{}

	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		rawKey = rsa
	default:
		return nil, fmt.Errorf("SSH: Unsupported key type %q", block.Type)
	}

	return newSignerFromKey(rawKey)
}

// Signer can create signatures that verify against a public key.
type Signer interface {
	Sign(data []byte) ([]byte, error)
}

func newSignerFromKey(k interface{}) (Signer, error) {
	var sshKey Signer

	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, fmt.Errorf("SSH: Unsupported key type %T", k)
	}

	return sshKey, nil
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d)
}
