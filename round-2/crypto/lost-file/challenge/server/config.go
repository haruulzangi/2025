package server

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"log"

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

func prepareServerConfig(oracleAddr string) (*ssh.ServerConfig, error) {
	serverKey, err := loadServerKey()
	if err != nil {
		return nil, fmt.Errorf("failed to load server key: %w", err)
	}

	config := &ssh.ServerConfig{
		PublicKeyAuthAlgorithms: []string{"ecdsa-sha2-nistp256"},
		BannerCallback: func(ctx ssh.ConnMetadata) string {
			return fmt.Sprintf("Welcome, %s! I don't know what you want, but please go to %s\n", ctx.User(), oracleAddr)
		},
		PublicKeyCallback: func(ctx ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			if subtle.ConstantTimeCompare(key.Marshal(), serverKey.Marshal()) == 1 {
				return &ssh.Permissions{}, nil
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
	return config, nil
}
