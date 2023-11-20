package utils

import (
	"math/rand"
	"time"
)

func GenerateKey() string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	word := ""
	for i := 0; i < 32; i++ {
		index := rand.Intn(len(characters))
		word += string(characters[index])
	}
	return word
}

func GenerateNonce() string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	word := ""
	for i := 0; i < 12; i++ {
		index := rand.Intn(len(characters))
		word += string(characters[index])
	}
	return word
}

func GenerateUserChatId() string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789<>/ç*&%$#@!~?"

	rand.Seed(time.Now().UnixNano())

	word := ""
	for i := 0; i < 50; i++ {
		index := rand.Intn(len(characters))
		word += string(characters[index])
	}
	return word
}

func GenerateToken() string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	rand.Seed(time.Now().UnixNano())

	word := ""
	for i := 0; i < 120; i++ {
		index := rand.Intn(len(characters))
		word += string(characters[index])
	}
	return word
}