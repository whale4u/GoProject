package main

import "fmt"

func init() {
	fmt.Println("init1:", a)
}

func init() {
	fmt.Println("init2:", b)
}

var a = 10

const b = 100

func main() {
	fmt.Println("main:", a)
}

//init1: 10
//init2: 10
//main: 10
