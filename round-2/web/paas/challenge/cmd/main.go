package main

import (
	"log"
	"net/http"
	"os"

	"github.com/haruulzangi/2025/challenges/round-2/web/paas/challenge/server"
)

func main() {
	server.DefinePaaSRoute()

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:8080"
	}
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
