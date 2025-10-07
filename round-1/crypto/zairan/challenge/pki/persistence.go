package pki

import (
	"crypto/ecdsa"
	"crypto/x509"
	"os"
	"path"
)

func loadCertificates(certFile string, keyFile string) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		return nil, nil, os.ErrNotExist
	}
	certPath := path.Join(dataDir, certFile)
	keyPath := path.Join(dataDir, keyFile)

	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		return nil, nil, os.ErrNotExist
	}
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return nil, nil, os.ErrNotExist
	}

	certPEM, err := os.ReadFile(certPath)
	if err != nil {
		return nil, nil, err
	}
	cert, err := x509.ParseCertificate(certPEM)
	if err != nil {
		return nil, nil, err
	}

	keyPEM, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, nil, err
	}
	priv, err := x509.ParseECPrivateKey(keyPEM)
	if err != nil {
		return nil, nil, err
	}
	return cert, priv, nil
}

func saveCertificates(certFile string, keyFile string, cert *x509.Certificate, key *ecdsa.PrivateKey) error {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		return nil
	}

	certPath := path.Join(dataDir, certFile)
	keyPath := path.Join(dataDir, keyFile)

	err := os.WriteFile(certPath, cert.Raw, 0644)
	if err != nil {
		return err
	}

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}
	err = os.WriteFile(keyPath, keyBytes, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(keyPath, keyBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}
