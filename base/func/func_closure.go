package main

import (
	"fmt"
)

//闭包（Closure）是匿名函数的一个特例。
// 当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()

	fmt.Println(a("World!"))
	fmt.Println(b("Niko!"))
}
