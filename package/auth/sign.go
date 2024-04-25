package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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
	base64DataString := base64.URLEncoding.EncodeToString(dataString)

	mac := hmac.New(sha256.New, []byte(SecretKey))
	signString := base64.URLEncoding.EncodeToString(mac.Sum([]byte(base64DataString)))
	return fmt.Sprintf("%s:%s:%s", SecretID, signString, base64DataString)
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
	validSignString := base64.URLEncoding.EncodeToString(mac.Sum([]byte(fragment[2])))
	if validSignString != fragment[1] {
		return false
	}

	tokenInfo, err := ParseToken(token)
	if err != nil {
		return false
	}

	// 过期了
	if tokenInfo.Data.ExpiresAt.Before(time.Now()) {
		return false
	}

	return true
}

type TokenInfo struct {
	SecretKey string
	Data      SignData
}

func ParseToken(token string) (*TokenInfo, error) {
	if separatorCount := strings.Count(token, ":"); separatorCount != 2 {
		return nil, fmt.Errorf("")
	}

	fragment := strings.Split(token, ":")
	SecretID := fragment[0]
	// signString := fragment[1]
	dataString := fragment[2]

	var tokenInfo TokenInfo
	var tokenData SignData
	decodedData, err := base64.URLEncoding.DecodeString(dataString)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(decodedData, &tokenData); err != nil {
		return nil, err
	}

	tokenInfo.SecretKey = SecretID
	tokenInfo.Data = tokenData
	return &tokenInfo, nil
}
