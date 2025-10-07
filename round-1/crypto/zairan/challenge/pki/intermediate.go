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

func generateIntermediateCertificate(rootCert *x509.Certificate, rootKey *ecdsa.PrivateKey) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	intermediateCert, intermediateKey, err := loadCertificates("intermediate.pem", "intermediate.key")
	if err == nil {
		return intermediateCert, intermediateKey, nil
	}

	intermediateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	intermediateTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: "Zairan Intermediate CA"},

		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(0, 6, 0),

		IsCA:           true,
		MaxPathLen:     0,
		MaxPathLenZero: true,

		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,

		AuthorityKeyId: rootCert.SubjectKeyId,
		SubjectKeyId:   CalculateSKI(&intermediateKey.PublicKey),
	}

	certDER, err := x509.CreateCertificate(
		rand.Reader,
		intermediateTemplate,
		rootCert,
		&intermediateKey.PublicKey,
		rootKey,
	)
	if err != nil {
		return nil, nil, err
	}
	intermediateCert, err = x509.ParseCertificate(certDER)
	if err != nil {
		return nil, nil, err
	}

	if err = saveCertificates("intermediate.pem", "intermediate.key", intermediateCert, intermediateKey); err != nil {
		log.Printf("Failed to save intermediate certificate: %v", err)
	}
	return intermediateCert, intermediateKey, nil
}
