package main

import (
	"fmt"
)

func main() {
	//创建名为boySalary的map，其中键是 string 类型，而值是 int 类型。
	boySalary := make(map[string]int)
	boySalary["bob"] = 1200

	//另一种创建方式
	girlSalary := map[string]int{
		"alice": 1300,
		//最后一个元素也要有逗号
		"susan": 1400,
	}

	fmt.Printf("map girlsalary's length is: %d\n", len(girlSalary))

	//遍历maps
	for key, value := range girlSalary {
		fmt.Printf("girl[%s] = %d\n", key, value)
	}

	//删除key
	fmt.Println("map before deletion", girlSalary)
	delete(girlSalary, "susan")
	fmt.Println("map after deletion", girlSalary)

	//判断map中是否包含某个key
	val, ok := girlSalary["susan"]
	if ok {
		fmt.Println("susan in girlSalary map")
		fmt.Println("Her salary is: ", val)
	} else {
		fmt.Println("susan is not in girlSalary map")
	}
}
