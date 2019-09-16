package middleware

import (
	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/pkg/errno"
	"tzh.com/web/pkg/token"
)

// 验证 JWT 的中间件
func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.GetHeader("Authorization")
		if _, err := token.Verify([]byte(t)); err != nil {
			handler.SendResponse(ctx, errno.New(errno.ErrTokenInvalid, err), nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
