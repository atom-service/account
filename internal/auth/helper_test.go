package auth

import "testing"

func TestIsGodSecret(t *testing.T) {
	if !IsGodSecret(GodSecretKey, GodSecretValue) {
		t.Error("Expected true, got false")
	}

	if IsGodSecret("mySecretKey", "mySecretValue") {
		t.Error("Expected true, got false")
	}
}
