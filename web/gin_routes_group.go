package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.POST("/submit", submit)
	}

	r.Run(":8080")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("Hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("Hello %s\n", name))
}

//测试
//http://127.0.0.1:8080/v1/login?name=niko
//curl http://127.0.0.1:8080/v1/submit -X POST
