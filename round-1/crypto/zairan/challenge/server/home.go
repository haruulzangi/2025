package server

import (
	"fmt"
	"net/http"
)

func DefineHomeRoute(oracleRoute string) {
	welcomeMessage := fmt.Sprintf("Welcome, stranger! We have /root.pem and /intermediate.pem routes for certificates.\n\nThe Zairan is waiting for you at %s…\nYou must please him by providing the client certificate if you know what i'm saying 😉", oracleRoute)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(welcomeMessage))
	})
}
