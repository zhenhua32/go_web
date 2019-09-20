package check

import (
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func PingServer(wait chan int) error {
	defer close(wait)
	time.Sleep(time.Second)
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/health")
		if err == nil && resp.StatusCode == 200 {
			wait <- 1
			return nil
		}

		logrus.Warnf("第 %d 失败, 一秒后重试", i)
		time.Sleep(time.Second)
	}

	return errors.New("不能连接到服务器, ping 失败")
}
