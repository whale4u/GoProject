package main

import (
	"fmt"
	"time"
)
/*
爬虫模块
解析模块
存储模块
*/

type MyURL struct {
	rc chan string
	wc chan string
}

func (mu *MyURL) SpiderURL() {
	startUrl := "https://www.baidu.com"
	mu.rc <- startUrl
}

func (mu *MyURL) Proccess() {
	//_ := <- mu.rc
	mu.wc <- "https://www.baidu.com"
}

func (mu *MyURL)WriteToDB() {
	fmt.Println(<-mu.wc)
}

func main() {
	myurl := &MyURL{
		rc: make(chan string),
		wc: make(chan string),
	}
	go myurl.SpiderURL()
	go myurl.Proccess()
	go myurl.WriteToDB()

	time.Sleep(1 * time.Second)
}