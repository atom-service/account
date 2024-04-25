package helper

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	// Test that the function returns a string of the correct length.
	str1 := GenerateRandomString(10, nil)
	if len(str1) != 10 {
		t.Errorf("Expected string length 10, got %d", len(str1))
	}

	// Test that the function returns a string containing only characters from the charset.
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for index := range str1 {
		if !containsChar(charset, str1[index]) {
			t.Errorf("String contains unexpected character %q", str1[index])
		}
	}

	// Test that the function returns a different string each time it is called.
	str2 := GenerateRandomString(10, nil)
	if str1 == str2 {
		t.Error("Expected different strings, but they were the same")
	}
}

func containsChar(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}
