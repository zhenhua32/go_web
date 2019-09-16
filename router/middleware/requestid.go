package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// 在请求头中设置 X-Request-Id
func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get("X-Request-Id")

		if requestId == "" {
			requestId = uuid.NewV4().String()
		}

		ctx.Set("X-Request-Id", requestId)

		ctx.Header("X-Request-Id", requestId)
		ctx.Next()
	}
}
