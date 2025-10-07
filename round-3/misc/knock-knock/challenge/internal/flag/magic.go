package flag

import (
	"encoding/base64"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

func deriveSignature(connMeta ssh.ConnMetadata) string {
	signature := base64.URLEncoding.EncodeToString(connMeta.SessionID())
	signature = strings.TrimRight(signature, "=")
	return signature
}

func GetFlag(connMeta ssh.ConnMetadata) string {
	flagTemplate := os.Getenv("FLAG")
	signature := deriveSignature(connMeta)
	signature_parts := []string{signature[:len(signature)/2], signature[len(signature)/2:]}
	return strings.Replace(strings.Replace(flagTemplate, "$1", signature_parts[0], 1), "$2", signature_parts[1], 1)
}
