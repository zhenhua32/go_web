package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tzh.com/web/handler/check"
	"tzh.com/web/router/middleware"
	"tzh.com/web/handler/user"
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

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
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
