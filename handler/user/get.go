package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/model"
	"tzh.com/web/pkg/errno"
)

func Get(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	user := model.UserModel{}

	if err := user.Fill(uint(userId)); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrFill, err), nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}

func GetByName(ctx *gin.Context) {
	username := ctx.Param("id")
	user, err := model.GetUserByName(username)
	if err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrFill, err), nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}
