package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	hash := hashAndSalt([]byte("123"))
	fmt.Println("hash过的密码", hash)
	pwdMatch := comparePasswords(hash, []byte("123"))
	if pwdMatch {
		fmt.Println("验证密码成功")
	}
}

// 加密密码
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// 验证密码
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
