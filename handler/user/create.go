package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tzh.com/web/handler"
	"tzh.com/web/pkg/errno"
)

func Create(ctx *gin.Context) {
	// 将 request body 绑定到一个结构体总
	var r CreateRequest

	if err := ctx.Bind(&r); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}
	logrus.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	admin := ctx.Param("username")
	logrus.Infof("URL param : %s", admin)

	desc := ctx.DefaultQuery("desc", "")
	logrus.Infof("URL query param des: %s", desc)

	contentType := ctx.GetHeader("Content-Type")
	logrus.Infof("Header Content-Type: %s", contentType)

	if r.Username == "" {
		err := errno.New(errno.ErrUserNotFound, fmt.Errorf("username 不在数据库中"))
		handler.SendResponse(ctx, err, nil)
		return
	}

	if r.Password == "" {
		err := fmt.Errorf("passwrod 是空的")
		handler.SendResponse(ctx, err, nil)
		return
	}

	resp := CreateResponse{
		Username: r.Username,
	}
	handler.SendResponse(ctx, nil, resp)
}
