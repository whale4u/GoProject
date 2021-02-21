package main

import (
	"fmt"
	"math"
)

func main() {
	//如果一个名字以大写字母开头，那么它就是已导出的。
	fmt.Println(math.Pi)
	//在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
	//fmt.Println(math.pi)
}
