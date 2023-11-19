package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(password string) string {
	hash_pass := sha256.Sum256([]byte(password))
	v := hex.EncodeToString(hash_pass[:])
	return v
}