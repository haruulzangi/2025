package server

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/haruulzangi/2025/challenges/round-1/crypto/zairan/challenge/flag"
)

func handleConnection(conn *tls.Conn, flagTemplate string) {
	defer conn.Close()
	if err := conn.Handshake(); err != nil {
		return
	}

	connState := conn.ConnectionState()
	if len(connState.PeerCertificates) > 0 {
		log.Printf("Client presented %d certificate(s), signature: %s", len(connState.PeerCertificates), flag.DeriveSignature(&connState))
		err := http.Serve(&singleConnListener{conn: conn}, CreateFlagMux(flag.GetFlag(flagTemplate, &connState)))
		if err != nil {
			log.Printf("Error serving HTTP over TLS: %v", err)
		}
		return
	}

	err := http.Serve(&singleConnListener{conn: conn}, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		log.Printf("Error serving HTTP over TLS: %v", err)
	}
}

func ServeTLS(flagTemplate string, rootCert *x509.Certificate, chain tls.Certificate, address string) {
	clientCAs := x509.NewCertPool()
	clientCAs.AddCert(rootCert)

	listener, err := tls.Listen("tcp", address, &tls.Config{
		Certificates: []tls.Certificate{chain},
		ClientCAs:    clientCAs,
		ClientAuth:   tls.VerifyClientCertIfGiven,
	})
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	log.Printf("Listening on %s…", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		tlsConn, ok := conn.(*tls.Conn)
		if !ok {
			log.Fatalf("Failed to assert connection as TLS :(")
		}
		go handleConnection(tlsConn, flagTemplate)
	}
}
