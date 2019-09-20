package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler"
	"tzh.com/web/model"
	"tzh.com/web/pkg/errno"
)

// @Summary 获取用户
// @Description 从数据库中获取用户信息
// @Tags user
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path uint64 true "user id in database"
// @Success 200 {object} model.UserModel "{"code":0,"message":"OK","data": {}}"
// @Router /user/{id} [get]
func Get(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	user := model.UserModel{}

	if err := user.Fill(uint(userId)); err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrFill, err), nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}

// @Summary 获取用户
// @Description 从数据库中获取用户信息
// @Tags user
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param username path string true "user's username in database"
// @Success 200 {object} model.UserModel "{"code":0,"message":"OK","data": {}}"
// @Router /user/{username}/byname [get]
func GetByName(ctx *gin.Context) {
	username := ctx.Param("id")
	user, err := model.GetUserByName(username)
	if err != nil {
		handler.SendResponse(ctx, errno.New(errno.ErrFill, err), nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}
