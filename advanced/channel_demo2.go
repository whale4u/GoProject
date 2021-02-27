package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func main() {
	done := make(chan bool)
	//
	go hello(done)
	//我们通过信道 done 接收数据。这一行代码发生了阻塞，除非有协程向 done 写入数据
	//<-done 这行代码通过信道done 接收数据，但并没有使用数据或者把数据存储到变量中。这完全是合法的。
	<-done
	fmt.Println("main function")
}
