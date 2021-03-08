package main

import (
	"fmt"
)

func main() {
	// 带参数的匿名函数
	f := func(args string) {
		fmt.Println(args)
	}

	f("Hello World!")
}
