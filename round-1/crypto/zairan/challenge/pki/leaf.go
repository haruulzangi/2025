package pki

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"time"
)

func SignLeafCertificate(dnsName string, intermediateCert *x509.Certificate, intermediateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (*x509.Certificate, error) {
	template := x509.Certificate{
		Subject:  pkix.Name{CommonName: dnsName},
		DNSNames: []string{dnsName},

		AuthorityKeyId: intermediateCert.SubjectKeyId,
		SubjectKeyId:   CalculateSKI(publicKey),

		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(30 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		IsCA:                  false,
		BasicConstraintsValid: true,
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, intermediateCert, publicKey, intermediateKey)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}
