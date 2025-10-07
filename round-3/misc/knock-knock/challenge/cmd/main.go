package main

import (
	"log"
	"os"

	"github.com/haruulzangi/2025/challenges/round-3/misc/knock-knock/challenge/internal/server"
)

func main() {
	flagTemplate := os.Getenv("FLAG")
	if flagTemplate == "" {
		log.Fatalf("FLAG environment variable was not provided")
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:2222"
	}
	if err := server.ListenAndServe(listenAddr); err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}
