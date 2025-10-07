package oracle

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"math/big"
)

type ECDSASignature struct {
	R, S *big.Int
}

func SignData(data []byte) ([]byte, error) {
	priv, err := loadPrivateKey()
	if err != nil {
		return nil, err
	}
	h := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, priv, h[:]) // ecdsa.SignASN1 seems to be using SHA512, we need SHA256 tho ¯\_(ツ)_/¯
	if err != nil {
		return nil, err
	}
	sig, err := asn1.Marshal(ECDSASignature{r, s})
	if err != nil {
		return nil, err
	}
	return sig, nil
}
