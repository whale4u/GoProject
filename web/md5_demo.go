package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte(str))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
