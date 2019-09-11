package util

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

// 生成短 id
func GenShortId() (string, error) {
	return shortid.Generate()
}

// 获取请求中的 X-Request-Id
func GetReqID(ctx *gin.Context) string {
	return ctx.GetHeader("X-Request-Id")
}

func TimeToStr(t *time.Time) string {
	if t == nil {
		return ""
	} else {
		return t.Format("2006-01-02 15:04:05")
	}
}
