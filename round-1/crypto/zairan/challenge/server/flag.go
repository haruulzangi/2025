package server

import (
	"net/http"
)

func CreateFlagMux(flag string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Yo, you did it! Now to go /flag and get your flag!!1"))
	})
	mux.HandleFunc("/flag", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(flag))
	})
	return mux
}
