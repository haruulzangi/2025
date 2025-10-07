package server

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha3"
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/haruulzangi/2025/challenges/round-3/crypto/futuristic-crypto/challenge/flag"
)

func FlagHandler(w http.ResponseWriter, r *http.Request) {
	mlkemKey, err := flag.GetMLKEMDecapsulationKey()
	if err != nil {
		http.Error(w, "Internal server error, please create a ticket in Discord.", http.StatusInternalServerError)
		log.Printf("Error retrieving MLKEM key: %v", err)
		return
	}

	alicePriv, err := flag.GetAlicePrivateKey()
	if err != nil {
		http.Error(w, "Internal server error, please create a ticket in Discord.", http.StatusInternalServerError)
		log.Printf("Error retrieving Alice's private key: %v", err)
		return
	}

	bobPriv, err := flag.GetBobPrivateKey()
	if err != nil {
		http.Error(w, "Internal server error, please create a ticket in Discord.", http.StatusInternalServerError)
		log.Printf("Error retrieving Bob's private key: %v", err)
		return
	}

	mlkeySharedSecret, kexCiphertext := mlkemKey.EncapsulationKey().Encapsulate()
	w.Write([]byte("Shared secret 1: " + base64.StdEncoding.EncodeToString(kexCiphertext) + "\n"))

	x25519SharedSecret, err := alicePriv.ECDH(bobPriv.PublicKey())
	if err != nil {
		http.Error(w, "Internal server error, please create a ticket in Discord.", http.StatusInternalServerError)
		log.Printf("Error performing X25519 ECDH: %v", err)
		return
	}
	w.Write([]byte("Performed a classic key exchange. Merging secrets with SHA3-256(secret1 || secret2)…\n"))

	hasher := sha3.New256()
	hasher.Write(mlkeySharedSecret)
	hasher.Write(x25519SharedSecret)
	sharedSecret := hasher.Sum(nil)
	w.Write([]byte("Futuristic key exchange complete, encrypting a flag for you with AES256-GCM :)\n"))

	flagTemplate := os.Getenv("FLAG")
	flag := flag.GetFlag(flagTemplate, sharedSecret)

	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		http.Error(w, "Internal server error, please create a ticket in Discord.", http.StatusInternalServerError)
		log.Printf("Error creating AES cipher: %v", err)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Internal server error, please create a ticket in Discord.", http.StatusInternalServerError)
		log.Printf("Error creating GCM: %v", err)
		return
	}

	ciphertext := gcm.Seal(nil, make([]byte, gcm.NonceSize()), []byte(flag), nil)
	w.Write([]byte("Here is your encrypted flag (base64): " + base64.StdEncoding.EncodeToString(ciphertext) + "\n"))
	w.Write([]byte("Good luck decrypting it!\n"))
}
