package user

import (
	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/pkg/errno"
	"tzh.com/web/service"
)

// @Summary 获取所有用户
// @Description 从数据库中获取所有用户的信息
// @Tags user
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param username query string false "part match username"
// @Param offset query int true "data offset" default(0)
// @Param limit query int true "data limit" default(10)
// @Success 200 {object} user.ListResponse "{"code":0,"message":"OK","data": {}}"
// @Router /user [get]
func List(ctx *gin.Context) {
	var r ListRequest
	if err := ctx.ShouldBindQuery(&r); err != nil {
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
