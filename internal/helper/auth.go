package helper

var (
	// 每次运行时总是生成不同的，以保证安全
	// 该 Secret 主要用于内部服务之间相互调用的情况
	// 用户的 secret 长度都是 64，这里的 128 是特殊的 
	GodSecretKey   = GenerateRandomString(128)
	GodSecretValue = GenerateRandomString(128)
)

// 判断是否是内部相互调用
func IsGodSecret(secretKey, secretValue string) bool {
	return GodSecretKey == secretKey && GodSecretValue == secretValue
}
