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
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	users, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	handler.SendResponse(ctx, nil, ListResponse{
		TotalCount: count,
		UserList:   users,
	})
}
