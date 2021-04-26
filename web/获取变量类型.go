package main

import (
	"fmt"
	"reflect"
)

func main() {
	v1 := "123456"
	v2 := 12

	//方法一
	fmt.Printf("v1 type:%T\n", v1)
	fmt.Printf("v2 type:%T\n", v2)

	//方法二
	fmt.Println("v1 type:", reflect.TypeOf(v1))
	fmt.Println("v2 type:", reflect.TypeOf(v2))
}
