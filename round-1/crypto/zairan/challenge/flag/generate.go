package flag

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"strings"
)

const SIGNATURE_TEMPLATE_VALUE = "SIGNATURE"

func merkleRoot(hashes [][]byte) []byte {
	if len(hashes) == 0 {
		return nil
	}
	if len(hashes) == 1 {
		return hashes[0]
	}

	var newLevel [][]byte
	for i := 0; i < len(hashes); i += 2 {
		if i+1 < len(hashes) {
			hash := sha256.Sum256(append(hashes[i], hashes[i+1]...))
			newLevel = append(newLevel, hash[:])
		} else {
			newLevel = append(newLevel, hashes[i])
		}
	}
	return merkleRoot(newLevel)
}

func DeriveSignature(connMeta *tls.ConnectionState) string {
	hashes := make([][]byte, 0)
	for _, peerCert := range connMeta.PeerCertificates {
		hashes = append(hashes, peerCert.Signature)
	}
	return strings.TrimRight(base64.URLEncoding.EncodeToString(merkleRoot(hashes)), "=")
}

func GetFlag(flagTemplate string, connMeta *tls.ConnectionState) string {
	signature := DeriveSignature(connMeta)
	return strings.Replace(flagTemplate, SIGNATURE_TEMPLATE_VALUE, signature, 1)
}
