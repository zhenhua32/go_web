package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tzh.com/web/handler"
	"tzh.com/web/pkg/errno"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 重写, 将同样的数据写一份保存到 body 中
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging 定义日志组件, 记录每一个请求
func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		ip := ctx.ClientIP()

		// 只记录特定的路由
		reg := regexp.MustCompile("(/v1/user|/login)")
		if !reg.MatchString(path) {
			return
		}

		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
		}
		// 读取后写回
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = blw

		start := time.Now()
		ctx.Next()
		// 计算延迟, 和 gin.Logger 的差距有点大
		// 这是因为 middleware 类似栈, 先进后出, ctx.Next() 是转折点
		// 所以 gin.Logger 放在最前, 记录总时长
		// Logging 放在最后, 记录实际运行的时间, 不包含其他中间件的耗时
		end := time.Now()
		latency := end.Sub(start)

		code, message := -1, ""
		var response handler.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			logrus.Errorf(
				"response body 不能被解析为 model.Response struct, body: `%s`, err: `%v`",
				blw.body.Bytes(),
				err,
			)
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}

		logrus.WithFields(logrus.Fields{
			"latency": fmt.Sprintf("%s", latency),
			"ip":      ip,
			"method":  method,
			"path":    path,
			"code":    code,
			"message": message,
		}).Info("记录请求")
	}
}
