package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tzh.com/web/pkg/errno"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 返回固定格式的响应结果
func SendResponse(ctx *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
