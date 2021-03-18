package main

import (
	"fmt"
)

//等价于
//type byte = uint8
//type rune = int32

//正如我们所说golang的string底层是一个byte数组实现，
// 中文字符串在unicode下占2个字节，在utf-8下占3个字节。
// 我们golang默认编码是utf-8，所以是占用3个字节。

func main() {
	var str = "小手25是什么"
	fmt.Println(str[:4]) //结果： 小�
	fmt.Println(len(str))
	fmt.Println(str[:8]) //结果：小手25
	//转换成切片
	s := []rune(str)
	fmt.Println(len(s))        //长度只有 7，每字汉字当一个字节
	fmt.Println(string(s[:4])) //结果：小手25

	tmpArr := map[int32]int{}
	tmpArr[1] = 666
	tmpArr[2] = 777

	if tmpArr[5] != 0 {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}

}
