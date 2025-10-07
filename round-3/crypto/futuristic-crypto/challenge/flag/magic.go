package flag

import (
	"encoding/base64"
	"log"
	"strings"
)

func GetFlag(flagTemplate string, sharedSecret []byte) string {
	flag := strings.Replace(flagTemplate, "$1", strings.TrimRight(base64.URLEncoding.EncodeToString(sharedSecret[:len(sharedSecret)/2]), "="), 1)
	flag = strings.Replace(flag, "$2", strings.TrimRight(base64.URLEncoding.EncodeToString(sharedSecret[len(sharedSecret)/2:]), "="), 1)
	log.Printf("Generated flag: %s", flag)
	return flag
}
