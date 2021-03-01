package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("start Goroutine")
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	// 3 个并发执行的 Go 协程（由 Go 主协程生成）
	for i := 0; i < no; i++ {
		wg.Add(i)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
