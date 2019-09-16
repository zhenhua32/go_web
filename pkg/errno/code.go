package errno

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/

var (
	OK = &Errno{Code: 0, Message: "OK"}

	// 系统错误, 前缀为 100
	InternalServerError = &Errno{Code: 10001, Message: "内部服务器错误"}
	ErrBind             = &Errno{Code: 10002, Message: "请求参数错误"}
	ErrTokenSign        = &Errno{Code: 10003, Message: "签名 jwt 时发生错误"}
	ErrEncrypt          = &Errno{Code: 10004, Message: "加密用户密码时发生错误"}

	// 数据库错误, 前缀为 201
	ErrDatabase = &Errno{Code: 20100, Message: "数据库错误"}
	ErrFill     = &Errno{Code: 20101, Message: "从数据库填充 struct 时发生错误"}

	// 认证错误, 前缀是 202
	ErrValidation   = &Errno{Code: 20201, Message: "验证失败"}
	ErrTokenInvalid = &Errno{Code: 20202, Message: "jwt 是无效的"}

	// 用户错误, 前缀为 203
	ErrUserNotFound      = &Errno{Code: 20301, Message: "用户没找到"}
	ErrPasswordIncorrect = &Errno{Code: 20302, Message: "密码错误"}
)
