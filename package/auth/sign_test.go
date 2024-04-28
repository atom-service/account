package auth

import (
	"testing"
	"testing/quick"
	"time"

	"github.com/atom-service/account/internal/helper"
)

func TestToken(t *testing.T) {
	config := &quick.Config{
		MaxCount: 100,
	}
	if err := quick.Check(func() bool {
		ak := helper.GenerateRandomString(64, nil)
		sk := helper.GenerateRandomString(64, nil)

		token := SignToken(ak, sk, SignData{
			ExpiresAt: time.Now().UTC().Add(time.Second),
		})

		tokenInfo, err := ParseToken(token)
		if err != nil {
			t.Errorf("ParseToken failed: %v", err)
			return false
		}

		if tokenInfo.SecretKey != ak {
			t.Errorf("ParseToken result are incorrect: %v", tokenInfo)
			return false
		}

		if ok := VerifyToken(ak, sk, token); !ok {
			t.Errorf("VerifyToken result are incorrect: %v", tokenInfo)
			return false
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}
}
