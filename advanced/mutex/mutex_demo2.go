package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	// Lock 和 Unlock之间不存在竞争条件。
	x += 1
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		//传递 Mutex 的地址很重要。
		// 如果传递的是 Mutex 的值，而非地址，那么每个协程都会得到 Mutex 的一份拷贝，
		// 竞态条件还是会发生。
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
