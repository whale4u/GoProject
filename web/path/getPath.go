package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	//"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	//i := strings.LastIndex(s, "\\")
	//path := string(s[0 : i+1])
	return s
}

func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return file
}

func main() {
	pa, _ := os.Getwd()
	path := getCurrentPath()
	filePath := CurrentFile()
	fmt.Println(pa)
	fmt.Println(path)
	fmt.Println(filePath)
}
