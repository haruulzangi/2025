package server

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/haruulzangi/2025/challenges/round-3/misc/knock-knock/challenge/internal/flag"
	"golang.org/x/crypto/ssh"
)

func addHostKey(config *ssh.ServerConfig) error {
	hostKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate ECDSA key: %v", err)
	}
	hostKeySigner, err := ssh.NewSignerFromKey(hostKey)
	if err != nil {
		log.Fatalf("Failed to create an SSH signer: %v", err)
	}
	config.AddHostKey(hostKeySigner)
	return nil
}

func prepareServerConfig() (*ssh.ServerConfig, error) {
	config := &ssh.ServerConfig{
		NoClientAuth: true,
		BannerCallback: func(ctx ssh.ConnMetadata) string {
			port, err := flag.SpawnFlagServer(ctx)
			if err != nil {
				log.Printf("Failed to spawn flag server: %v", err)
				return "Sorry, something went wrong. Please reconnect.\n"
			}
			return fmt.Sprintf("Welcome, %s! I don't know what you want, but please go to tcp://localhost:%d\n", ctx.User(), port)
		},
	}

	addHostKey(config)
	return config, nil
}
