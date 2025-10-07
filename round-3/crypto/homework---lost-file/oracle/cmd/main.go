package main

import (
	"net/http"
	"os"

	"github.com/haruulzangi/2025/challenges/round-2/crypto/lost-file/oracle/server"
)

func main() {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		panic("DATA_DIR environment variable is not set")
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:8080"
	}

	server.DefineOracleRoute()
	http.ListenAndServe(listenAddr, nil)
}
