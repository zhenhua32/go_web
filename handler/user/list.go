package user

import (
	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/pkg/errno"
	"tzh.com/web/service"
)

func List(ctx *gin.Context) {
	var r ListRequest
	if err := ctx.Bind(&r); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrBind, err), nil)
		return
	}

	users, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		// 对于有多种类型的错误, 直接用 err 就行了
		handler.SendResponse(ctx, err, nil)
		return
	}

	handler.SendResponse(ctx, nil, ListResponse{
		TotalCount: count,
		UserList:   users,
	})
}
