package main

//文件操作示例

import (
	"fmt"
	"os"
)

func main() {
	//新建文件夹
	os.Mkdir("test", 0755)
	os.MkdirAll("a/b/c", 0755)

	//删除文件夹
	err := os.Remove("test")
	if err != nil {
		fmt.Println("error:", err)
	}

	//递归删除文件夹
	os.RemoveAll("a")
}
