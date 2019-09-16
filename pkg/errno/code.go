package errno

/*
错误码设计
第一位表示错位分类, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务器错误"}
	ErrBind             = &Errno{Code: 10002, Message: "请求参数错误"}

	ErrValidation = &Errno{Code: 20001, Message: "验证失败"}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库错误"}
	ErrToken      = &Errno{Code: 20003, Message: "签名 JSON web token 时发生错误"}
	ErrFill       = &Errno{Code: 20004, Message: "从数据库填充 struct 时发生错误"}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "加密用户密码时发生错误"}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "用户没找到"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "token 是无效的"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码错误"}
)
