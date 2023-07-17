package utils

import (
	"encoding/base64"
)

func Encrypt(text string, key string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(key + text)), nil
}

func GenerateKey(token string) string {
	return token[:16]
}
