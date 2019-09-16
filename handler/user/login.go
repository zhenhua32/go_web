package user

import (
	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/model"
	"tzh.com/web/pkg/errno"
	"tzh.com/web/pkg/token"
)

func Login(ctx *gin.Context) {
	var u model.UserModel
	// 应该使用 ShouldBindJSON, 以便使用自定义的 handler.SendResponse
	if err := ctx.ShouldBindJSON(&u); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrBind, err), nil)
		return
	}

	user, err := model.GetUserByName(u.Username)
	if err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrDatabase, err), nil)
		return
	}

	if err := user.Compare(u.Password); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrPasswordIncorrect, err), nil)
		return
	}

	// 签发 token
	t, err := token.Sign(user.ID, user.Username)
	if err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrTokenSign, err), nil)
		return
	}
	handler.SendResponse(ctx, nil, model.Token{Token: t})
}
