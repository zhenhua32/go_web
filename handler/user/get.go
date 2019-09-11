package user

import (
	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/model"
)

func Get(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := model.GetUserByName(username)
	if err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}
