package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `title`
	Content string `string`
	Tag     string `tag`
}

// 定义接收数据的结构体
type PP struct {
	Params []Article
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	r.POST("/loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var json PP
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		//if json.User != "root" || json.Pssword != "admin" {
		//	c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		//	return
		//}
		//POSTMAN请求 {"params": [{"title":"A", "content":"B", "tag":"xxx"}]}
		//fmt.Printf("%+v",json)
		fmt.Println(json.Params)
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8080")
}
