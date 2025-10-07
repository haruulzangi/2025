package main

import (
	"log"
	"net/http"
	"os"

	"github.com/haruulzangi/2025/challenges/round-3/crypto/futuristic-crypto/challenge/flag"
	"github.com/haruulzangi/2025/challenges/round-3/crypto/futuristic-crypto/challenge/server"
)

func main() {
	if err := flag.EnsureKeysExist(); err != nil {
		log.Fatalf("Failed to ensure keys exist: %v", err)
	}

	fs := http.FileServer(http.Dir(flag.GetDataPath()))
	http.Handle("/keys/", http.StripPrefix("/keys", fs))

	flagTemplate := os.Getenv("FLAG")
	if flagTemplate == "" {
		log.Fatal("FLAG environment variable is not set")
	}

	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/key-exchange/flag", server.FlagHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
