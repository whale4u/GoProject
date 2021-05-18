package main

import "fmt"

//声明接口
type Animal interface {
	GetAge() int32
	GetType() string
}

type Dog struct {
	Age  int32
	Type string
}

//dog实现了animal接口
func (a Dog) GetAge() int32 {
	return a.Age
}

func (a Dog) GetType() string {
	return a.Type
}

func main() {
	animal := Dog{Age: 20, Type: "DOG"}
	fmt.Printf("%s max age is: %d", animal.GetType(), animal.GetAge())

}
