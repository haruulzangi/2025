package server

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"

	"github.com/haruulzangi/2025/challenges/round-2/crypto/lost-file/oracle/oracle"
)

const ORACLE_ROUTE = "/zairan"

func DefineOracleRoute() string {
	http.HandleFunc(ORACLE_ROUTE, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method, expected POST with data", http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get("Content-Type") != "application/octet-stream+base64" {
			http.Error(w, "Invalid Content-Type, expected application/octet-stream+base64", http.StatusBadRequest)
			return
		}
		if r.ContentLength <= 0 {
			http.Error(w, "Invalid Content-Length, expected non-zero length", http.StatusBadRequest)
			return
		}

		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
			http.Error(w, "Failed to read request body", 499) // https://http.cat/status/499
			return
		}

		rawData, err := base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			http.Error(w, "Failed to decode base64 data", http.StatusBadRequest)
			return
		}
		signature, err := oracle.SignData(rawData)
		if err != nil {
			log.Printf("Failed to sign data: %v", err)
			http.Error(w, "Failed to sign data. Please create a ticket in Discord.", 500)
			return
		}

		w.Header().Set("Content-Type", "application/octet-stream+base64")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(base64.StdEncoding.EncodeToString(signature)))
	})
	return ORACLE_ROUTE
}
