package flag

import (
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/ssh"
)

const SIGNATURE_TEMPLATE_VALUE = "SIGNATURE"

func DeriveSignature(connMeta ssh.ConnMetadata) string {
	signature := base64.URLEncoding.EncodeToString(connMeta.SessionID())
	signature = strings.TrimRight(signature, "=")
	return signature
}

func GetFlag(flagTemplate string, connMeta ssh.ConnMetadata) string {
	signature := DeriveSignature(connMeta)
	return strings.Replace(flagTemplate, SIGNATURE_TEMPLATE_VALUE, signature, 1)
}
