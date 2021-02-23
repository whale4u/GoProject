package main

import "fmt"

//可变参数函数的工作原理是把可变参数转换为一个新的切片。
func find(num int, nums ...int) {
	fmt.Printf("type of num is %T\n", nums)
	found := false

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
	//有一个可以直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。
	// 如果这样做，切片将直接传入函数，不再创建新的切片
	//直接传入切片，不在创建新切片
	nums := []int{89, 90, 95}
	find(89, nums...)
}
