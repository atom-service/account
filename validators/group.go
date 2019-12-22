package validators

import "regexp"

var (
	groupNamePattern        = regexp.MustCompile(`^[a-zA-Z0-9-_]{4,128}$`) // 组名称
	groupCategoryPattern       = regexp.MustCompile(`^[a-zA-Z0-9-_]{4,128}$`) // 组名称
	groupStatePattern       = regexp.MustCompile(`^[a-zA-Z0-9-_]{4,128}$`) // 组名称
	groupDescriptionPattern = regexp.MustCompile(`^.{4,512}$`)             // 组名称
)

// GroupName 组名
func GroupName(value string) (pass bool, msg string) {
	if !groupNamePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字、中划线、下划线，长度 4 - 128 位"
	}
	return true, ""
}

// GroupCategory 组类
func GroupCategory(value string) (pass bool, msg string) {
	if !groupCategoryPattern.MatchString(value) {
		return false, "仅支持大小写字母、数字、中划线、下划线，长度 4 - 128 位"
	}
	return true, ""
}

// GroupState 组状态
func GroupState(value string) (pass bool, msg string) {
	if !groupStatePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字、中划线、下划线，长度 4 - 128 位"
	}
	return true, ""
}

// GroupDescription 组说明
func GroupDescription(value string) (pass bool, msg string) {
	if !groupDescriptionPattern.MatchString(value) {
		return false, "支持任意字符、长度 4 - 512 位"
	}
	return true, ""
}
