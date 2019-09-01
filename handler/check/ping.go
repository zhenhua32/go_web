package check

import (
	"errors"
	"log"
	"net/http"
	"time"
)

func PingServer() error {
	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("第一次失败, 一秒后重试")
		time.Sleep(time.Second)
	}

	return errors.New("不能连接到服务器, ping 失败")
}
