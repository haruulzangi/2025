package server

import (
	"fmt"
	"net/http"

	"github.com/haruulzangi/2025/challenges/round-2/web/paas/challenge/sandbox"
)

var content = []byte(`<html>
<head><title>Ping-as-a-service</title></head>
<body>
<h1>Welcome to the Ping-as-a-Service!</h1>
	<form method="POST">
		<label for="url">Enter URL to ping:</label><br>
		<input type="text" id="url" name="url" size="50"><br><br>
		<input type="submit" value="Ping">
	</form>
</body>
</html>`)

func DefinePaaSRoute() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Add("Content-Type", "text/html")
			w.Write(content)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		url := r.Form.Get("url")
		if url == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		exitCode := sandbox.RunWithTimeout(fmt.Sprintf("ping -c 4 %s", url))
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html><body>"))
		w.Write([]byte("<h2>Ping Result</h2>"))
		w.Write([]byte(fmt.Sprintf("<p>Status: %d</p>", exitCode)))
		w.Write([]byte("<a href=\"/\">Back</a>"))
		w.Write([]byte("</body></html>"))
	})
}
