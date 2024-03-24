package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		index := newRand.Intn(len(charset))
		randomString[i] = charset[index]
	}

	return string(randomString)
}
