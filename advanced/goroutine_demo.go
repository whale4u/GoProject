package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched ()
		//表示让 CPU 把时间片让给别人，下次某个时候继续恢复执行该 goroutine。
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	// 开启一个新的goroutines执行
	go say("world")
	// 当前goroutines执行
	say("hello")
}
