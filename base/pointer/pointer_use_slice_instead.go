package main

import (
	"fmt"
)

func modify(arr []int) {
	arr[0] = 90
}

func main() {
	array := [3]int{89, 90, 91}
	//使用切片传递数组
	modify(array[:])
	fmt.Println(array)
}
