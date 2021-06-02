package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	//返回 reflect.Type 具体类型
	t := reflect.TypeOf(q)
	//该类型的特定类别
	k := t.Kind()
	//返回 reflect.Value 具体值
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t.String())
	fmt.Println("Kind ", k)
	fmt.Println("Value ", v)

	fmt.Println("iterate struct")
	//NumField返回结构体中的key数量，Field(i)返回i的reflect.value
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, v.Field(i))
	}
}

func main() {
	o := order{
		ordId:      456,
		customerId: 78,
	}
	createQuery(o)
}
