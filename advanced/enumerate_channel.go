package main

import "fmt"

func producer(chn1 chan int) {
	for i := 0; i < 10; i++ {
		chn1 <- i
	}
	//关闭信道
	close(chn1)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	//一种写法
	//for {
	//	v, ok := <-ch
	//	if ok == false {
	//		break
	//	}
	//	fmt.Println("Received: ", v, ok)
	//}

	//另一种写法
	for v := range ch {
		fmt.Println("Received: ", v)
	}
}
