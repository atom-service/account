package http

// HttpCode 是 Http 状态码的枚举类型
type HttpCode int

// 定义常见的 Http 状态码
const (
	HttpCodeOK           HttpCode = 200
	HttpCodeBadRequest   HttpCode = 400
	HttpCodeUnauthorized HttpCode = 401
	HttpCodeNotFound     HttpCode = 404
	HttpCodeServerError  HttpCode = 500
)

// String 返回 HttpCode 的字符串表示
func (c HttpCode) String() string {
	switch c {
	case HttpCodeOK:
		return "OK"
	case HttpCodeBadRequest:
		return "请求错误"
	case HttpCodeUnauthorized:
		return "未登录"
	case HttpCodeNotFound:
		return "请求的资源不存在"
	case HttpCodeServerError:
		return "内部服务错误"
	default:
		return "未知错误"
	}
}

type HttpResult struct {
	Code    HttpCode
	Message *string
	Data    interface{}
}

func CreateHttpResult(code HttpCode, data interface{}, message *string) *HttpResult {
	result := &HttpResult{
		Code:    code,
		Data:    data,
		Message: message,
	}

	if result.Message == nil {
		normal := code.String()
		result.Message = &normal
	}

	return result
}
