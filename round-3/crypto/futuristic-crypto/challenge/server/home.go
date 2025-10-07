package server

import (
	"fmt"
	"net/http"

	"github.com/haruulzangi/2025/challenges/round-3/crypto/futuristic-crypto/challenge/flag"
)

var data = fmt.Sprintf(`Welcome to the Futuristic Crypto Challenge!
Here, you'll explore cutting-edge cryptographic techniques and solve intriguing puzzles.

Available endpoints:
/keys/%s
/keys/%s
/keys/%s

Go grab your flag at: /key-exchange/flag!
Good luck and have fun!
`, flag.MLKEMKeyFile, flag.AliceKeyFile, flag.BobKeyFile)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(data))
}
