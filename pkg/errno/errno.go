package errno

import "fmt"

// 定义错误码
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// 定义错误
type Err struct {
	Code    int
	Message string // 展示给用户看的
	Errord  error  // 保存内部错误信息
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Errord)
}

// 使用 错误码 和 error 创建新的 错误
func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Errord:  err,
	}
}

func (err *Err) Add(message string) *Err {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) *Err {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// 解码错误, 获取 Code 和 Message
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
