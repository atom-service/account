package provider

import "regexp"

var (
	nicknamePattern = regexp.MustCompile(`^.{2,128}$`)
	usernamePattern = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`)
	passwordPattern = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`)
)
