package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tzh.com/web/handler"
	"tzh.com/web/model"
	"tzh.com/web/pkg/errno"
	"tzh.com/web/util"
)

// 完整更新, 所有的字段都应该传递
func Save(ctx *gin.Context) {
	logrus.WithField(
		"X-Request-Id", util.GetReqID(ctx),
	).Info("用户更新函数被调用")
	userId, _ := strconv.Atoi(ctx.Param("id"))

	var u model.UserModel
	if err := ctx.Bind(&u); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrBind, err), nil)
		return
	}

	u.ID = uint(userId)

	// 校验数据
	if err := u.Validate(); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrValidation, err), nil)
		return
	}

	// 加密用户密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrEncrypt, err), nil)
		return
	}

	// 更新数据库
	if err := u.Save(); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrDatabase, err), nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}

// 选择更新, 只更新传递的字段
func Update(ctx *gin.Context) {
	logrus.WithField(
		"X-Request-Id", util.GetReqID(ctx),
	).Info("用户更新函数被调用")
	userId, _ := strconv.Atoi(ctx.Param("id"))

	var data map[string]interface{}
	if err := ctx.Bind(&data); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrBind, err), nil)
		return
	}

	// 验证字段, 并加密密码
	if err := model.ValidateAndUpdateUser(&data); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrValidation, err), nil)
		return
	}

	// 更新数据库
	user := &model.UserModel{}
	user.ID = uint(userId)
	if err := user.Update(data); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrDatabase, err), nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}
