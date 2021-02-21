package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "/Users/whale4u/Code/GoProject/base/test.txt"
	//该方法打开一个名称为 name 的文件，但是是只读方式，内部实现其实调用了 OpenFile。
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	defer fl.Close()
	//创建动态数组，只能用于 slice，map，channel 三种类型
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
