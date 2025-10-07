package pki

import (
	"crypto"
	"crypto/sha1"
	"crypto/x509"
)

func CalculateSKI(publicKey crypto.PublicKey) []byte {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	ski := sha1.Sum(pubKeyBytes)
	return ski[:]
}
