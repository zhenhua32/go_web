package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler/check"
	"tzh.com/web/handler/user"
	"tzh.com/web/router/middleware"
)

// 载入中间件
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache())
	g.Use(middleware.Options())
	g.Use(middleware.Secure())
	g.Use(mw...)

	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "incorrect api router")
	})

	g.POST("/login", user.Login)

	u := g.Group("/v1/user")
	u.Use(middleware.AuthJWT()) // 添加认证
	{
		u.GET("", user.List)
		u.POST("", user.Create)
		u.PUT("/:id", user.Save)
		u.DELETE("/:d", user.Delete)
		u.GET("/:id", user.Get)
		u.PATCH("/:id", user.Update)
		// 设计得不太好, 只能用 id 作为占位符
		u.GET("/:id/byname", user.GetByName)
	}

	checkRoute := g.Group("/check")
	{
		checkRoute.GET("/health", check.HealthCheck)
		checkRoute.GET("/disk", check.DiskCheck)
		checkRoute.GET("/cpu", check.CPUCheck)
		checkRoute.GET("/memory", check.MemoryCheck)
	}

	return g

}
