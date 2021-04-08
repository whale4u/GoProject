package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

//curl请求：必须是json请求
//curl --location --request POST 'http://127.0.0.1:8088/login' \
//--header 'Content-Type: application/json' \
//--data-raw '{"username":"admin", "password":"pass"}'

func main() {
	r := gin.Default()
	r.POST("/login", Login)
	//r.GET("/logout", getting)
	//监听端口默认为8080
	r.Run(":8088")
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		fmt.Println(err)
		return
	}
	//find user with username
	//user, err := models.UserRepo.FindByID(1)
	//compare the user from the request, with the one we defined:
	if u.UserName != "admin" || u.Password != "pass" {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token := CreateToken(u.UserName, u.Password)
	//if err != nil {
	//	c.JSON(http.StatusUnprocessableEntity, err.Error())
	//	return
	//}

	// save token to redis
	//saveErr := servers.HttpServer.RD.CreateAuth(user.ID, ts)
	//if saveErr != nil {
	//	c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	//}

	c.JSON(http.StatusOK, token)
}

func CreateToken(userName string, password string) string {

	hmacSampleSecret := []byte("my_secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userName,
		"password": password,
		"nbf":      time.Date(2019, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)
	fmt.Println("===name====")
	fmt.Println(VerifyToken(tokenString))
	return tokenString

}

func VerifyToken(tokenString string) (bool, string) {
	var name string
	hmacSampleSecret := []byte("my_secret_key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		//}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		log.Print("请检查 token 参数！")
		return false, name
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], claims["nbf"])
		name = claims["username"].(string)
	} else {
		fmt.Println(err)
		return false, name
	}

	return true, name
}
