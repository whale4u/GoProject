package main

import (
	"fmt"
)

// 定义函数类型add
type add func(a int, b int) int

func main() {
	// 声明a为add类型
	var a add
	a = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum ", s)
}
