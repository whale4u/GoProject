package main

import (
	"fmt"
)

func main() {
	//先声明，后续再初始化
	var a chan int

	//另一种初始化方法
	b := make(chan int)

	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
		fmt.Printf("Type of b is %T\n", b)
	}
}
