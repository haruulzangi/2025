package pki

import (
	"crypto/ecdsa"
	"crypto/x509"
)

func GenerateCertificateBundle() (rootCert *x509.Certificate, intermediateCert *x509.Certificate, intermediateKey *ecdsa.PrivateKey) {
	rootCertObj, rootKey, err := generateRootCertificate()
	if err != nil {
		return
	}
	intermediateCert, intermediateKey, err = generateIntermediateCertificate(rootCertObj, rootKey)
	if err != nil {
		return
	}
	return rootCertObj, intermediateCert, intermediateKey
}
