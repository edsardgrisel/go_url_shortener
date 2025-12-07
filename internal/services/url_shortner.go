package services

import (
	"crypto/sha1"
	"encoding/hex"
)

// "crypto/sha1"

func Url_shortener(url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	hash := hasher.Sum(nil)
	return (hex.EncodeToString(hash)[:6])
}
