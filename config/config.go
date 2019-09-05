package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
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
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置已经被改变: %s", e.Name)
		// time.Sleep(time.Second)
		if err := viper.ReadInConfig(); err != nil {
			return
		}
		change <- 1
	})
}
