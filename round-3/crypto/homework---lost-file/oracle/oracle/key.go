package oracle

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path"
)

func loadPrivateKey() (*ecdsa.PrivateKey, error) {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		return nil, os.ErrNotExist
	}
	keyPEM, err := os.ReadFile(path.Join(dataDir, "key.pem"))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(keyPEM)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing EC private key")
	}
	priv, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}
