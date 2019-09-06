package config

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

// 读取配置
func (c *Config) InitConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")

	// 从环境变量总读取
	viper.AutomaticEnv()
	viper.SetEnvPrefix("web")
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	return viper.ReadInConfig()
}

// 监控配置改动
func (c *Config) WatchConfig(change chan int) {
	viper.WatchConfig()
	// TODO: 这个会触发两次, 考虑使用限流模式
	// https://github.com/gohugoio/hugo/blob/master/watcher/batcher.go
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("配置已经被改变: %s", e.Name)

		// time.Sleep(time.Second)
		if err := viper.ReadInConfig(); err != nil {
			return
		}
		change <- 1
	})
}

// 初始化日志
func (c *Config) InitLog() {
	// log.use_json
	if viper.GetBool("log.use_json") {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	// log.logger_level
	switch viper.GetString("log.logger_level") {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	}

	// log.logger_file
	logger_file := viper.GetString("log.logger_file")
	os.MkdirAll(filepath.Dir(logger_file), os.ModePerm)

	file, err := os.OpenFile(logger_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(file)
	}

	// log.gin_file & log.gin_console
	gin_file := viper.GetString("log.gin_file")
	os.MkdirAll(filepath.Dir(gin_file), os.ModePerm)

	file, err = os.OpenFile(gin_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		if viper.GetBool("log.gin_console") {
			gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
		} else {
			gin.DefaultWriter = io.MultiWriter(file)
		}
	}

	// default
	logrus.SetReportCaller(true)
}
