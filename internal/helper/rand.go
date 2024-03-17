package helper

import (
	"math/rand"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		index := rand.Intn(len(charset))
		randomString[i] = charset[index]
	}

	return string(randomString)
}
