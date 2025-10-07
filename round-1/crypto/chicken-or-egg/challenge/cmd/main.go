package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	flag "github.com/haruulzangi/2025/challenges/round-1/crypto/chicken-or-egg/flag"
	"golang.org/x/crypto/ssh"
)

func generateClientKey() (sshPublicKey string, pemKey string) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate ECDSA key: %v", err)
	}
	sshKey, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		log.Fatalf("Failed to create SSH public key: %v", err)
	}
	sshPublicKey = string(sshKey.Marshal())

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		log.Fatalf("Failed to marshal EC key: %v", err)
	}
	pemKey = string(pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes}))
	return sshPublicKey, pemKey
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:2222"
	}

	flagTemplate := os.Getenv("FLAG")
	if flagTemplate == "" {
		log.Fatalf("FLAG environment variable was not provided")
	}
	if !strings.Contains(flagTemplate, flag.SIGNATURE_TEMPLATE_VALUE) {
		log.Fatalf("Invalid flag format, expected %s to be present", flag.SIGNATURE_TEMPLATE_VALUE)
	}

	authKeys := make(map[string]string)
	config := &ssh.ServerConfig{
		PublicKeyAuthAlgorithms: []string{"ecdsa-sha2-nistp256"},
		BannerCallback: func(ctx ssh.ConnMetadata) string {
			sshPublicKey, pemKey := generateClientKey()
			authKeys[hex.EncodeToString(ctx.SessionID())] = string(sshPublicKey)
			return fmt.Sprintf("Welcome, %s! Chicken or egg?\n\n%s", ctx.User(), pemKey)
		},
		PublicKeyCallback: func(ctx ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			if authKeys[hex.EncodeToString(ctx.SessionID())] == string(key.Marshal()) {
				log.Printf("Successful login from %s with signature %s", ctx.RemoteAddr(), flag.DeriveSignature(ctx))
				return nil, nil
			}
			return nil, fmt.Errorf("unknown key")
		},
	}

	hostKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate ECDSA key: %v", err)
	}
	hostKeySigner, err := ssh.NewSignerFromKey(hostKey)
	if err != nil {
		log.Fatalf("Failed to create an SSH signer: %v", err)
	}
	config.AddHostKey(hostKeySigner)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %s", listenAddr, err)
	}
	log.Printf("Listening on %s...", listenAddr)

	for {
		nConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept incoming connection: %s", err)
			continue
		}
		go func(nConn net.Conn) {
			sshConn, chans, reqs, err := ssh.NewServerConn(nConn, config)
			if err != nil {
				nConn.Close()
				return
			}
			defer sshConn.Close()
			go ssh.DiscardRequests(reqs)
			for newChannel := range chans {
				if newChannel.ChannelType() != "session" {
					newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
					continue
				}
				channel, requests, err := newChannel.Accept()
				if err != nil {
					continue
				}
				go ssh.DiscardRequests(requests)

				message := fmt.Sprintln(flag.GetFlag(flagTemplate, sshConn))
				io.WriteString(channel, message)
				channel.Close()
			}
		}(nConn)
	}
}
