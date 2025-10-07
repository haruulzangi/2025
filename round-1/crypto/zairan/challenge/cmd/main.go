package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"strings"

	"github.com/haruulzangi/2025/challenges/round-1/crypto/zairan/challenge/flag"
	"github.com/haruulzangi/2025/challenges/round-1/crypto/zairan/challenge/pki"
	"github.com/haruulzangi/2025/challenges/round-1/crypto/zairan/challenge/server"
)

func buildCertChain(certs []*x509.Certificate, leafPrivateKey *ecdsa.PrivateKey) tls.Certificate {
	leaf := certs[0]
	intermediate := certs[1]
	root := certs[2]

	return tls.Certificate{
		Certificate: [][]byte{leaf.Raw, intermediate.Raw, root.Raw},
		PrivateKey:  leafPrivateKey,
	}
}

func main() {
	flagTemplate := os.Getenv("FLAG")
	if flagTemplate == "" {
		log.Fatalf("FLAG environment variable was not provided")
	}
	if !strings.Contains(flagTemplate, flag.SIGNATURE_TEMPLATE_VALUE) {
		log.Fatalf("Invalid flag format, expected %s to be present", flag.SIGNATURE_TEMPLATE_VALUE)
	}

	rootCert, intermediateCert, intermediateKey := pki.GenerateCertificateBundle()
	leafKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	dnsName := os.Getenv("DNS_NAME")
	if dnsName == "" {
		dnsName = "zairan.challenge.haruulzangi.mn"
	}
	leafCert, err := pki.SignLeafCertificate(dnsName, intermediateCert, intermediateKey, &leafKey.PublicKey)
	if err != nil {
		panic(err)
	}
	server.DefinePKIRoutes(rootCert, intermediateCert)
	oracleRoute := server.DefineOracleRoute(intermediateCert, intermediateKey)
	server.DefineHomeRoute(oracleRoute)

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:8443"
	}
	chain := buildCertChain([]*x509.Certificate{leafCert, intermediateCert, rootCert}, leafKey)
	server.ServeTLS(flagTemplate, rootCert, chain, listenAddr)
}
