package check

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func PingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Printf("第 %d 失败, 一秒后重试", i)
		time.Sleep(time.Second)
	}

	return errors.New("不能连接到服务器, ping 失败")
}
