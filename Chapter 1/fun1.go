package main

import "fmt"

//func add(x int, y int) int  {
//	return x + y
//}

//也可以这样写
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(44, 55))
}
