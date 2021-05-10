package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// 例如@every 1h表示每小时触发一次，@every 1m2s表示每隔 1 分 2 秒触发一次。
	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})

	// 启动一个新的 goroutine 做循环检测
	c.Start()
	// 防止主 goroutine 退出
	time.Sleep(time.Second * 5)
}
