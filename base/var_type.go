package main

import "fmt"

func main() {
	// 声明并赋值方式一
	var flagA bool //默认值为false
	flagA = true
	// 也可以向下面这样写
	//var flagA bool = true

	// 声明并赋值方式一
	flagB := true

	// 声明常量
	const PI = 3.14159

	fmt.Println("flagA is: ", flagA)
	fmt.Println("flagB is: ", flagB)
}
