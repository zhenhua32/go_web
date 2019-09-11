package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/model"
)

func Get(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	user := model.UserModel{}
	user.ID = uint(userId)

	if err := user.Fill(); err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}

func GetByName(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := model.GetUserByName(username)
	if err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}
