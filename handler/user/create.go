package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tzh.com/web/pkg/errno"
)

func Create(ctx *gin.Context) {
	// 将 request body 绑定到一个结构体总
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err = ctx.Bind(&r); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    errno.ErrBind.Code,
			"message": errno.ErrBind.Message,
		})
		return
	}
	logrus.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username 不在数据库中"))
		logrus.Error(err, "; 发生了一个错误")
	} else if r.Password == "" {
		err = fmt.Errorf("passwrod 是空的")
	}

	code, message := errno.DecodeErr(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})

}
