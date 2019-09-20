package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/model"
	"tzh.com/web/pkg/errno"
)

// @Summary 删除用户
// @Description 在数据库中标记用户为删除状态
// @Tags user
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path uint64 true "user id in database"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data": null}"
// @Router /user/{id} [delete]
func Delete(ctx *gin.Context) {
	// 将文本转换为字符串
	userId, _ := strconv.Atoi(ctx.Param("id"))
	if err := model.DeleteUser(uint(userId)); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrDatabase, err), nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}
