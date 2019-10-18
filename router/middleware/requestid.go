package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestID 在请求头中设置 X-Request-Id
func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.Request.Header.Get("X-Request-Id")

		if requestID == "" {
			requestID = uuid.NewV4().String()
		}

		ctx.Set("X-Request-Id", requestID)

		ctx.Header("X-Request-Id", requestID)
		ctx.Next()
	}
}
