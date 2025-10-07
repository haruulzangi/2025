package pki

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
	"time"
)

func generateRootCertificate() (*x509.Certificate, *ecdsa.PrivateKey, error) {
	cert, priv, err := loadCertificates("root.pem", "root.key")
	if err == nil {
		return cert, priv, nil
	}

	priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	ski := CalculateSKI(&priv.PublicKey)
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Zairan CA"},
		},
		SubjectKeyId:   ski,
		AuthorityKeyId: ski,

		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),

		KeyUsage: x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		IsCA:     true,

		BasicConstraintsValid: true,
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}
	cert, err = x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, nil, err
	}

	if err = saveCertificates("root.pem", "root.key", cert, priv); err != nil {
		log.Printf("Failed to save root certificate: %v", err)
	}
	return cert, priv, nil
}
