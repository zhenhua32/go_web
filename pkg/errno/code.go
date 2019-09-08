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
	ErrBind             = &Errno{Code: 10002, Message: "绑定请求体到 stuct 时发生错误"}

	// user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "用户不存在"}
)
