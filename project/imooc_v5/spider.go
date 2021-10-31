package main

import (
	"fmt"
	"time"
	//"regexp"
	"mvdan.cc/xurls/v2"
)
/*
爬虫模块
解析模块
存储模块
*/

type MyURL struct {
	rc chan string
	wc chan []string
}

func (mu *MyURL) SpiderURL() {
	startUrl := "https://www.baidu.com https://www.sina.com"
	mu.rc <- startUrl
}

func (mu *MyURL) Proccess() {
	url := <- mu.rc
	xurlsRelaxed := xurls.Relaxed()
	output := xurlsRelaxed.FindAllString(url, 2)
	//fmt.Println(output)
	//reg := regexp.MustCompile(`[a-zA-z]+://[^\s]*`)
	//-1表示返回所有匹配的值，1表示返回1个匹配的值，以此类推
	//fmt.Println(reg.FindAllString(url, -1))
	mu.wc <- output
}

func (mu *MyURL)WriteToDB() {
	fmt.Println(<-mu.wc)
}

func main() {
	myurl := &MyURL{
		rc: make(chan string),
		wc: make(chan []string),
	}
	go myurl.SpiderURL()
	go myurl.Proccess()
	go myurl.WriteToDB()

	time.Sleep(1 * time.Second)
}