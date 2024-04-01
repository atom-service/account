package validator

import (
	"regexp"
	"strings"
)

type Validator func(string) bool

func WithMinLength(minLength int) Validator {
	return func(value string) bool {
		return len(value) >= minLength
	}
}

func WithMaxLength(maxLength int) Validator {
	return func(value string) bool {
		return len(value) <= maxLength
	}
}

func WithAllowedText(allowedText string) Validator {
	return func(value string) bool {
		for _, c := range value {
			if !strings.Contains(allowedText, string(c)) {
				return false
			}
		}
		return true
	}
}

func WithEmail() Validator {
	return func(value string) bool {
		pattern := `^[\w\.-]+@[\w\.-]+\.\w+$`
		match, _ := regexp.MatchString(pattern, value)
		return match
	}
}

func WithURL() Validator {
	return func(value string) bool {
		pattern := `^(http|https)://[\w\.-]+\.\w+$`
		match, _ := regexp.MatchString(pattern, value)
		return match
	}
}

func ValidateString(input string, options ...Validator) bool {
	for _, option := range options {
		if !option(input) {
			return false
		}
	}

	return true
}
