package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string
	Order string
}

func main() {
	// 注意这里byte后面是反引号
	var jsonBlob = []byte(` [
	{"Name": "Plat", "Order": "ppp"},
	{"Name": "Quoll", "Order": "Das"}
	]`)

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}

	//%v值的默认格式表示。当输出结构体时，扩展标志（%+v）会添加字段名
	fmt.Printf("%+v", animals)
}
