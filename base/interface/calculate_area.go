package main

import "fmt"

//声明Caculator接口
type Caculator interface {
	CaculateArea() int
	CaculatorVol() int
}

type Triangle struct {
	length int
	width  int
}

type Rectangle struct {
	length int
	width  int
}

func (t Triangle) CaculateArea() int {
	return t.width * t.length / 2
}

func (r Rectangle) CaculateArea() int {
	return r.width * r.length
}

func main() {
	t := Triangle{
		length: 5,
		width:  3, //这里也需要逗号
	}
	fmt.Println(t.CaculateArea())
}
