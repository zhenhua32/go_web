package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/pkg/errno"
	"tzh.com/web/pkg/token"
)

// AuthJWT 验证 JWT 的中间件
func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		headerList := strings.Split(header, " ")
		if len(headerList) != 2 {
			err := errors.New("无法解析 Authorization 字段")
			handler.SendResponse(ctx, errno.New(errno.ErrTokenInvalid, err), nil)
			ctx.Abort()
			return
		}
		t := headerList[0]
		content := headerList[1]
		if t != "Bearer" {
			err := errors.New("认证类型错误, 当前只支持 Bearer")
			handler.SendResponse(ctx, errno.New(errno.ErrTokenInvalid, err), nil)
			ctx.Abort()
			return
		}
		if _, err := token.Verify([]byte(content)); err != nil {
			handler.SendResponse(ctx, errno.New(errno.ErrTokenInvalid, err), nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
