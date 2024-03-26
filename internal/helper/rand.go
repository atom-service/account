package helper

import (
	"math/rand"
	"time"
)

var defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(length int, charset *string) string {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	if charset == nil {
		charset = &defaultCharset
	}

	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		index := newRand.Intn(len(*charset))
		randomString[i] = (*charset)[index]
	}

	return string(randomString)
}
