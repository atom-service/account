package validators

import "regexp"

var (
	labelNamePattern  = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`) // 标签名称
	labelClassPattern = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`) // 标签名称
	labelStatePattern = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`) // 标签名称
	labelValuePattern = regexp.MustCompile(`^[a-zA-Z0-9]{6,128}$`) // 标签名称
)

func LabelName(value string) (pass bool, msg string) {
	if !labelNamePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}

func LabelClass(value string) (pass bool, msg string) {
	if !labelClassPattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}

func LabelState(value string) (pass bool, msg string) {
	if !labelStatePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}

func LabelValue(value string) (pass bool, msg string) {
	if !labelValuePattern.MatchString(value) {
		return false, "仅支持大小写字母、数字，长度 2- 128 位"
	}
	return true, ""
}
