package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"tzh.com/web/handler"
	"tzh.com/web/model"
	"tzh.com/web/pkg/auth"
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
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	u.ID = uint(userId)

	// 校验数据
	if err := u.Validate(); err != nil {
		handler.SendResponse(ctx, errno.ErrValidation, nil)
		return
	}

	// 加密用户密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(ctx, errno.ErrEncrypt, nil)
		return
	}

	// 更新数据库
	if err := u.Save(); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}

// 另一种方式
func Update(ctx *gin.Context) {
	logrus.WithField(
		"X-Request-Id", util.GetReqID(ctx),
	).Info("用户更新函数被调用")
	userId, _ := strconv.Atoi(ctx.Param("id"))

	var data map[string]interface{}
	if err := ctx.Bind(&data); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	// 验证, 不太好做
	var userV *model.UserModel
	if err := mapstructure.Decode(&data, userV); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}
	if err := userV.Validate(); err != nil {
		handler.SendResponse(ctx, errno.ErrValidation, nil)
		return
	}

	// 加密用户密码
	if password, ok := data["password"]; ok {
		newPassword, err := auth.Encrypt(password.(string))
		if err == nil {
			data["password"] = newPassword
		}
	}

	// 更新数据库
	user := &model.UserModel{}
	user.ID = uint(userId)
	if err := user.Update(data); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}
