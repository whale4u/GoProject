package main

import (
	"fmt"
)

func change(val *int) {
	// 在 change 函数内使用解引用，修改了 a 的值
	*val = 55
}
func main() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	// 向函数 change 传递了指针变量 b，而 b 存储了 a 的地址
	change(b)
	fmt.Println("value of a after function call is", a)
}
