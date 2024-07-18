package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// encrypt the password using SHA-256
func EncodePassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashedPassword := hash.Sum(nil)
	hashString := hex.EncodeToString(hashedPassword)

	return hashString
}
