package main

import (
	"fmt"
)

//高阶函数条件：
//1、接收一个或者多个函数作为参数
//2、返回值是一个函数

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	s := simple()
	fmt.Println(s(60, 7))
}
