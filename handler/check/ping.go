package check

import (
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func PingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		logrus.Warnf("第 %d 失败, 一秒后重试", i)
		time.Sleep(time.Second)
	}

	return errors.New("不能连接到服务器, ping 失败")
}
