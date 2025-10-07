package server

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path"

	"golang.org/x/crypto/ssh"
)

func loadServerKey() (ssh.PublicKey, error) {
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
	pub, err := ssh.NewPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, err
	}
	return pub, nil
}
