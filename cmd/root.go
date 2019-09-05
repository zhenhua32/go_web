package cmd

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"tzh.com/web/config"
	"tzh.com/web/handler/check"
	"tzh.com/web/router"
)

var cfgFile string
var configChange = make(chan int, 1)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server is a simple restful api server",
	Long: `server is a simple restful api server
	use help get more ifo`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(args)
		runServer()
	},
}

// 定义 rootCmd 命令的执行
func runServer() {
	// 设置运行模式
	gin.SetMode(viper.GetString("runmode"))

	// 初始化空的服务器
	app := gin.New()
	// 保存中间件
	middlewares := []gin.HandlerFunc{}

	// 路由
	router.Load(
		app,
		middlewares...,
	)

	// 检查服务器正常启动
	go func() {
		if err := check.PingServer(); err != nil {
			log.Fatal("服务器没有响应:", err)
		}
		log.Printf("服务器正常启动")
	}()

	// 服务器裕兴的地址和端口
	addr := viper.GetString("addr")
	log.Printf("启动服务器在 http address: %s", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: app,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待配置改变, 然后重启
	<-configChange
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	runServer()
}

// 初始化, 设置 flag 等
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./conf/config.yaml)")
}

// 初始化配置
func initConfig() {
	c := config.Config{
		Name: cfgFile,
	}

	if err := c.InitConfig(); err != nil {
		panic(err)
	}
	log.Printf("载入配置成功")
	c.WatchConfig(configChange)
}

// 包装了 rootCmd.Execute()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
