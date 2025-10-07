package flag

import (
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/ssh"
)

func DeriveSignature(connMeta ssh.ConnMetadata) string {
	signature := base64.URLEncoding.EncodeToString(connMeta.SessionID())
	signature = strings.TrimRight(signature, "=")
	return signature
}

func GetFlag(flagTemplate string, connMeta ssh.ConnMetadata) string {
	signature := DeriveSignature(connMeta)
	flag := strings.Replace(flagTemplate, "$1", signature[:len(signature)/2], 1)
	return strings.Replace(flag, "$2", signature[len(signature)/2:], 1)
}
