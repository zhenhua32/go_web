package util

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

// GenShortID 生成短 id
func GenShortID() (string, error) {
	return shortid.Generate()
}

// GetReqID 获取请求中的 X-Request-Id
func GetReqID(ctx *gin.Context) string {
	return ctx.GetHeader("X-Request-Id")
}

// TimeToStr 将日期转换为字符串
func TimeToStr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// GetTag 获取结构体中的 tag
func GetTag(structed interface{}, fieldname string, tagname string) (string, bool) {
	t := reflect.TypeOf(structed)
	field, ok := t.FieldByName(fieldname)
	if !ok {
		return "", false
	}
	tag, ok := field.Tag.Lookup(tagname)
	return tag, ok
}
