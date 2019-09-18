package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"tzh.com/web/handler/check"
	"tzh.com/web/model"
	"tzh.com/web/router"
	"tzh.com/web/router/middleware"
)

// 定义 rootCmd 命令的执行
func runServer() {
	wait := make(chan int, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	// 设置运行模式
	gin.SetMode(viper.GetString("runmode"))

	// 初始化空的服务器
	app := gin.New()
	// 保存中间件
	middlewares := []gin.HandlerFunc{
		middleware.RequestId(),
		middleware.Logging(),
	}

	// 路由
	router.Load(
		app,
		middlewares...,
	)

	// 检查服务器正常启动
	go func() {
		if err := check.PingServer(wait); err != nil {
			logrus.Fatal("服务器没有响应:", err)
		}
		logrus.Info("服务器正常启动")
	}()

	// 服务器的地址和端口
	addr := viper.GetString("addr")

	srv := &http.Server{
		Addr:    addr,
		Handler: app,
	}
	// 启动服务
	go func() {
		logrus.Infof("启动服务器在 http address: %s", addr)
		srv.Addr = addr
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen on http: %s\n", err)
		}
	}()
	/*
		自签名的语句
		MSYS_NO_PATHCONV=1 openssl req -new -nodes -x509
		-out conf/server.crt -keyout conf/server.key -days 3650
		-subj "/C=CN/ST=SH/L=SH/O=CoolCat/OU=CoolCat Software/CN=127.0.0.1/emailAddress=coolcat@qq.com"
	*/
	// 启动 https 服务
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	addrTLS := viper.GetString("tls.addr")
	if cert != "" && key != "" {
		go func() {
			// 等待 http 服务正常启动
			<-wait
			logrus.Infof("启动服务器在 https address: %s", addrTLS)
			srv.Addr = addrTLS
			if err := srv.ListenAndServeTLS(cert, key); err != nil && err != http.ErrServerClosed {
				logrus.Fatalf("listen on https: %s\n", err)
			}
		}()
	}

	// 等待配置改变 或者 收到退出的信号
	select {
	case <-configChange:
		if err := srv.Shutdown(context.Background()); err != nil {
			logrus.Fatal("Server Shutdown:", err)
		}
		defer runServer()
		return
	case <-quit:
		logrus.Info("收到退出的信号")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logrus.Fatal("Server Shutdown: ", err)
		}
		logrus.Println("Server exiting")
		return
	}
}
