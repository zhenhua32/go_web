package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tzh.com/web/handler"
	"tzh.com/web/model"
	"tzh.com/web/pkg/errno"
	"tzh.com/web/util"
)

// 创建一个新的用户帐号
func Create(ctx *gin.Context) {
	logrus.WithField(
		"X-Request-Id", util.GetReqID(ctx),
	).Info("用户创建函数被调用")
	// 将 request body 绑定到一个结构体总
	var r CreateRequest
	if err := ctx.Bind(&r); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}
	logrus.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 验证结构
	if err := u.Validate(); err != nil {
		handler.SendResponse(ctx, errno.ErrValidation, nil)
		return
	}

	// 加密密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(ctx, errno.ErrEncrypt, nil)
		return
	}

	// 插入用户到数据库中
	if err := u.Create(); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Username: r.Username,
	}
	handler.SendResponse(ctx, nil, resp)
}
