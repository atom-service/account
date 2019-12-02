package validators

import "regexp"

var (
	nicknamePattern = regexp.MustCompile(`^.{2,128}$`)
	usernamePattern = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`)
	passwordPattern = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`)

	groupNamePattern        = regexp.MustCompile(`^.{2,128}$`)           // 组名称
	groupClassPattern       = regexp.MustCompile(`^.{2,128}$`)           // 组名称
	groupStatePattern       = regexp.MustCompile(`^.{2,128}$`)           // 组名称
	groupDescriptionPattern = regexp.MustCompile(`^.{2,128}$`)           // 组名称
	labelNamePattern        = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`) // 标签名称
	labelClassPattern       = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`) // 标签名称
	labelStatePattern       = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`) // 标签名称
	labelValuePattern       = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`) // 标签名称
)

func nicknameValidator(value string) (pass bool, msg string) {
	if !nicknamePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}

func usernameValidator(value string) (pass bool, msg string) {
	if !nicknamePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}

func passwordValidator(value string) (pass bool, msg string) {
	if !nicknamePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}
