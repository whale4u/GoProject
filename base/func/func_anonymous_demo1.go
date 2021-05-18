package main

import (
	"fmt"
)

func main() {
	// 带参数匿名函数
	f := func() {
		fmt.Println("I'm Anonymous Function")
	}
	f()
	fmt.Println("%T", f)
}
