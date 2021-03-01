package main

import (
	"fmt"
)

func main() {
	//3是指的是信道容量
	ch := make(chan string, 3)
	ch <- "niko"
	ch <- "lucky"

	fmt.Println("capacity is ", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value ", <-ch)
	fmt.Println("new length is", len(ch))
}
