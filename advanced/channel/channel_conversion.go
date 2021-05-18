package main

import (
	"fmt"
)

//函数 sendData 里的参数 sendch chan<- int 把 cha1 转换为一个唯送信道
func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)
}
