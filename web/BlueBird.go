package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"reflect"
)

// 定义一个全局对象db
var db *sql.DB

type Article struct {
	Title   string `title, required`
	Content string `string, required`
	Tag     string `tag, required`
}

// 定义接收数据的结构体
type Note struct {
	Params []Article
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insertMultiRow(note Note) {

	// 插入数据
	stmt, err := db.Prepare("INSERT into article(title, content, tag) value (?, ?, ?)")
	checkErr(err)

	var list [3]string
	//遍历切片中的结构体元素
	for _, v := range note.Params {
		value := reflect.ValueOf(v)
		for i := 0; i < value.NumField(); i++ {
			//fmt.Printf("Field %d: %v\n", i, value.Field(i))
			list[i] = value.Field(i).String()
		}
		//fmt.Println(list)
		_, err := stmt.Exec(list[0], list[1], list[2])
		checkErr(err)
	}

	//非常重要！！！
	defer stmt.Close()
}

func main() {

	initDB()

	r := gin.Default()

	//POSTMAN请求：{"params": [{"title":"A", "content":"B", "tag":"xxx"}, {"title":"B1","content":"B2","tag":"xxx1"}]}
	r.POST("/loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var note Note
		// 将request的body中的数据，自动按照note格式解析到结构体
		if err := c.ShouldBindJSON(&note); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//fmt.Printf("%+v",note)
		//fmt.Println(note.Params)
		c.JSON(http.StatusOK, gin.H{"status": "200"})

		//遍历切片中的结构体元素
		//for _, v := range note.Params {
		//	value := reflect.ValueOf(v)
		//	for i := 0; i < value.NumField(); i++ {
		//			fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//		}
		//	}

		insertMultiRow(note)
	})
	r.Run(":8080")
}
