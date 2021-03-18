package main

import (
	"fmt"
)

//例子2：在N个字符串找到找没有重复字符，且字符串总长度最长的那个

func main() {
	ss := []string{
		"ado",
		"duzhenxun",
		"小手25是什么",
		"来个长点的字符串，微信号5552123",
	}

	maxLenStr := ""
	for i := 0; i < len(ss); i++ {
		var repeat bool
		//声明并赋值map，只是赋值为空，这样就不用make了。
		tmpArr := map[int32]int{}
		//[]rune(ss[i])转换成切片
		for k, v := range []rune(ss[i]) {
			//判断v字符在map中是否存在，存在时tmpArr[v] != 0
			if tmpArr[v] != 0 && len(tmpArr) > 0 {
				repeat = true //有重复
				break
			}
			fmt.Println(k, v)
			//如果不存在，把字符v作为key加入map中
			tmpArr[v] = k
		}
		//没有重复找最长的
		if !repeat && len(ss[i]) > len(maxLenStr) {
			maxLenStr = ss[i]
		}
	}
	fmt.Println("无重复最长的是：", maxLenStr)

	//结果：小手25是什么
}
