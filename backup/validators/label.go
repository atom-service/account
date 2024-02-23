package validators

import "regexp"

var (
	labelNamePattern     = regexp.MustCompile(`^[a-zA-Z0-9-_]{4,128}$`) // 标签名称
	labelCategoryPattern = regexp.MustCompile(`^[a-zA-Z0-9-_]{4,128}$`) // 标签类型
	labelStatePattern    = regexp.MustCompile(`^[a-zA-Z0-9-_]{4,128}$`) // 标签状态
	labelValuePattern    = regexp.MustCompile(`^.{4,512}$`)             // 标签值 任意字符
)

// LabelName 标签名称
func LabelName(value string) (pass bool, msg string) {
	if labelNamePattern.MatchString(value) {
		return true, ""
	}
	return false, "标签名仅支持大小写字母、数字、中划线、下划线，长度 4 - 128 位"
}

// LabelCategory 类型
func LabelCategory(value string) (pass bool, msg string) {
	if !labelCategoryPattern.MatchString(value) {
		return false, "标签类别仅支持大小写字母、数字、中划线、下划线，长度 4 - 128 位"
	}
	return true, ""
}

// LabelState 状态
func LabelState(value string) (pass bool, msg string) {
	if !labelStatePattern.MatchString(value) {
		return false, "标签状态仅支持大小写字母、数字、中划线、下划线，长度 4 - 128 位"
	}
	return true, ""
}

// LabelValue 值
func LabelValue(value string) (pass bool, msg string) {
	if !labelValuePattern.MatchString(value) {
		return false, "任意字符，长度 4 - 512 位"
	}
	return true, ""
}
