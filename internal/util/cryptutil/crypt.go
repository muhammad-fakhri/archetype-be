package cryptutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// EncryptWithSHA256 encrypt by using SHA256
func EncryptWithSHA256(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
