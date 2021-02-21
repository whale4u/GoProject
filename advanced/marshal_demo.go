package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID int
	// 注意json前面是反引号，字段标记：这些标记信息通过反射接口可见，并参与结构体的类型标识，但在其他情况下被忽略。
	Name   string `json:"name"`
	Colors []string
	note   string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Redis",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		//只有首字母大写的成员才可以序列化为json，所以note不会被序列化
		note: "this is note",
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
