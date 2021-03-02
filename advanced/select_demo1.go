package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	//select 会选择首先响应的服务器，而忽略其它的响应。
	// 使用这种方法，我们可以向多个服务器发送请求，并给用户返回最快的响应了。:）
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
		//default:
		//	fmt.Println("no value received！")
		// 加了会立即执行到default！！！
	}
}
