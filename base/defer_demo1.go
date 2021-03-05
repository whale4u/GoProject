package main

// defer 不仅限于函数[4]的调用，调用方法[5]也是合法的

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	p := person{
		firstName: "Niko",
		lastName:  "Buck",
	}
	defer p.fullName()
	fmt.Printf("Hello ")
}
