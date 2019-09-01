package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler/check"
	"tzh.com/web/router"
)

func main() {
	// 初始化空的服务器
	app := gin.New()
	// 保存中间件
	middlewares := []gin.HandlerFunc{}

	// 路由
	router.Load(
		app,
		middlewares...,
	)

	go func() {
		if err := check.PingServer(); err != nil {
			log.Fatal("服务器没有响应", err)
		}
		log.Printf("服务器正常启动")
	}()

	log.Printf("启动服务器在 http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", app).Error())
}
