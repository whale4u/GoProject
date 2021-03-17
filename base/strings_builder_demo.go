package main

//Go 语言中，字符串是只读的，也就意味着每次修改操作都会创建一个新的字符串。
//如果需要拼接多次，应使用 strings.Builder，最小化内存拷贝次数。

import (
	"fmt"
	"strings"
)

func main() {
	var str strings.Builder

	for i := 0; i < 1000; i++ {
		str.WriteString("a")
	}
	fmt.Println(str.String())
}
