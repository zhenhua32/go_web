package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/model"
	"tzh.com/web/pkg/errno"
)

func Delete(ctx *gin.Context) {
	// 将文本转换为字符串
	userId, _ := strconv.Atoi(ctx.Param("id"))
	if err := model.DeleteUser(uint(userId)); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrDatabase, err), nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}
