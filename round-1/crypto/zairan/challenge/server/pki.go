package server

import (
	"crypto/x509"
	"encoding/pem"
	"net/http"
)

func convertCertToPEM(der []byte) []byte {
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: der,
	}
	return pem.EncodeToMemory(block)
}

func DefinePKIRoutes(rootCert *x509.Certificate, intermediateCert *x509.Certificate) {
	rootPEM := convertCertToPEM(rootCert.Raw)
	intermediatePEM := convertCertToPEM(intermediateCert.Raw)
	http.HandleFunc("/root.pem", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-pem-file")
		w.Write(rootPEM)
	})
	http.HandleFunc("/intermediate.pem", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-pem-file")
		w.Write(intermediatePEM)
	})
}
