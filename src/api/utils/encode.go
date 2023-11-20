package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func Encrypt(k, n, c string) string {
	key := []byte(k)
	plaintext := []byte(c)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte(n)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(k, n, c string) string {
	key := []byte(k)
	ciphertext, _ := hex.DecodeString(c)
	nonce := []byte(n)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}