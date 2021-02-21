package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "write.txt"
	//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是 0666 的文件，返回的文件对象是可读写的。
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		//fout.WriteString("Just a test! \r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	//删除文件
	os.Remove(userFile)
}
