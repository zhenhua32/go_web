package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"tzh.com/web/config"
)

var cfgFile string
var configChange = make(chan int, 1)

// 定义主命令 server
var rootCmd = &cobra.Command{
	Use:   "web",
	Short: "web is a simple restful api server",
	Long: `web is a simple restful api server
	use help get more ifo`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("启动参数: ", args)
		runServer()
	},
}

// 初始化, 设置 flag 等
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./conf/config.yaml", "config file (default: ./conf/config.yaml)")
	rootCmd.AddCommand(versionCmd)
}

// 初始化配置
func initConfig() {
	c := config.Config{
		Name: cfgFile,
	}

	if err := c.InitConfig(); err != nil {
		panic(err)
	}
	c.InitLog()
	logrus.Info("载入配置成功")
	c.WatchConfig(configChange)
}

// 包装了 rootCmd.Execute()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
