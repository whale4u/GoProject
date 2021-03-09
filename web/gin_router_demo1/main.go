package main

import (
	"fmt"
)

func main() {
	r := setupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

//curl http://127.0.0.1:8080/hello
//https://www.liwenzhou.com/posts/Go/gin_routes_registry/
