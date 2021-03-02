package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	//由于缓冲信道的容量为 1，所以任何其他协程试图写入该信道时，都会发生阻塞，
	// 直到 x 增加后，信道的值才会被读取（第 10 行）
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
