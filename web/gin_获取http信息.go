package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		fmt.Println(c.Request.URL.String())
		fmt.Println(c.Request.Method)
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.Run()
}
