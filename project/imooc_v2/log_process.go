package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc    chan []byte
	wc    chan string
	read  Reader
	write Writer
}

type ReadFromFile struct {
	path string // 读取文件路径
}

type WriteToInfluxDB struct {
	influxDBDsn string // influx data source
}

func (r *ReadFromFile) Read(rc chan []byte) {
	//	读取模块
	//line := "message"
	//rc <- line
	// 打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}
	// 从文件末尾开始读取文件内容
	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		//去掉换行符
		rc <- line[:len(line)-1]
	}

}

func (w *WriteToInfluxDB) Write(wc chan string) {
	//	写入模块
	//fmt.Println(<-wc)
	for v := range wc {
		fmt.Println(v)
	}
}

func (l *LogProcess) Process() {
	//	解析模块
	//data := <-l.rc
	//l.wc <- strings.ToUpper(string(data))
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

func main() {
	r := &ReadFromFile{
		path: "/Users/whale4u/Code/GoProject/project/imooc_v2/access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDsn: "username&password..",
	}

	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan string),
		read:  r,
		write: w,
	}
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(30 * time.Second) //等待协程执行完成
}
