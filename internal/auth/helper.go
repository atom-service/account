package auth

import "github.com/atom-service/account/internal/helper"

var (
	// 每次运行时总是生成不同的，以保证安全
	// 该 Secret 主要用于内部服务之间相互调用的情况
	// 用户的 secret 长度都是 64，这里的 128 是特殊的
	GodSecretKey   = helper.GenerateRandomString(128, nil)
	GodSecretValue = helper.GenerateRandomString(128, nil)
)

// 判断是否是内部相互调用
// 用于内部测试用途的 secret
func IsGodSecret(secretKey, secretValue string) bool {
	return GodSecretKey == secretKey && GodSecretValue == secretValue
}
