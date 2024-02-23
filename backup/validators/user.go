package validators

import "regexp"

var (
	nicknamePattern = regexp.MustCompile(`^.{4,128}$`)
	usernamePattern = regexp.MustCompile(`^[0-9A-Za-z_]{6,128}$`)
	passwordPattern = regexp.MustCompile(`^[A-Za-z0-9!#\$%&\(\)\*\+,-\./:;\?@\[\]]{6,128}$`) // !#$%&()*+,-./:;?@[]
)

// Nickname 昵称检查
func Nickname(value string) (pass bool, msg string) {
	if nicknamePattern.MatchString(value) {
		return true, ""
	}
	return false, "昵称支持任意字符，长度 4 - 128 位"
}

// Username 用户名
func Username(value string) (pass bool, msg string) {
	if usernamePattern.MatchString(value) {
		return true, ""
	}
	return false, "用户名仅支持大小写字母、数字、下划线，长度 6 - 128 位"
}

// Password 密码
func Password(value string) (pass bool, msg string) {
	if passwordPattern.MatchString(value) {
		return true, ""
	}
	return false, "用户密码仅支持大小写字母、数字、下划线以及 !#$%&()*+,-./:;?@[] 等字符, 长度 6 - 128 位"

}
