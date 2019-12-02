package validators

import "regexp"

var (
	groupNamePattern        = regexp.MustCompile(`^.{2,128}$`) // 组名称
	groupClassPattern       = regexp.MustCompile(`^.{2,128}$`) // 组名称
	groupStatePattern       = regexp.MustCompile(`^.{2,128}$`) // 组名称
	groupDescriptionPattern = regexp.MustCompile(`^.{2,128}$`) // 组名称
)

func GroupName(value string) (pass bool, msg string) {
	if !groupNamePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}

func GroupClass(value string) (pass bool, msg string) {
	if !groupClassPattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}

func GroupState(value string) (pass bool, msg string) {
	if !groupStatePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}

func GroupDescription(value string) (pass bool, msg string) {
	if !groupDescriptionPattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}
