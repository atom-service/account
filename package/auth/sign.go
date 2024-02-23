package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type SignData struct {
	ExpiresAt time.Time `json:"expiresAt"`
}

func SignToken(SecretID string, SecretKey string, data SignData) string {
	dataString, _ := json.Marshal(data)
	mac := hmac.New(sha256.New, []byte(SecretKey))
	signString := string(mac.Sum([]byte(dataString)))
	return fmt.Sprintf("%s:%s:%s", SecretID, signString, dataString)
}

func VerifyToken(SecretID string, SecretKey string, token string) bool {
	if separatorCount := strings.Count(token, ":"); separatorCount != 2 {
		return false
	}

	fragment := strings.Split(token, ":")
	if fragment[0] != SecretID {
		return false
	}

	mac := hmac.New(sha256.New, []byte(SecretKey))
	validSignString := string(mac.Sum([]byte(fragment[2])))
	if validSignString != fragment[1] {
		return false
	}

	return true
}

type TokenInfo struct {
	SecretID string
	Data SignData
}

func ParseToken(token string) (*TokenInfo, error) {
	if separatorCount := strings.Count(token, ":"); separatorCount != 2 {
		return nil, fmt.Errorf("")
	}

	return nil, nil
}
