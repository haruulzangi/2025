package main

import (
	"log"
	"os"

	"github.com/haruulzangi/2025/challenges/round-2/crypto/lost-file/challenge/server"
)

func main() {
	oracleAddr := os.Getenv("ORACLE_ADDR")
	if oracleAddr == "" {
		log.Fatalf("ORACLE_ADDR environment variable was not provided")
	}

	flagTemplate := os.Getenv("FLAG")
	if flagTemplate == "" {
		log.Fatalf("FLAG environment variable was not provided")
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:2222"
	}
	if err := server.ListenAndServe(flagTemplate, listenAddr, oracleAddr); err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}
