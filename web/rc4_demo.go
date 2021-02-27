package main

import (
	"crypto/rc4"
	"fmt"
)

//是一种流加密算法，密钥长度可变。它加解密使用相同的密钥，因此也属于对称加密算法。
// RC4是有线等效加密（WEP）中采用的加密算法，也曾经是TLS可采用的算法之一。

func main() {

	// 密钥
	key := []byte("abcdefg")
	// 要加密的源数据
	str := []byte("this is my test!")

	// 加密方式1：加密/解密后的数据单独存放
	{
		// 加密操作
		dest1 := make([]byte, len(str))
		fmt.Printf("方法1加密前:%s \n", str)
		cipher1, _ := rc4.NewCipher(key)
		cipher1.XORKeyStream(dest1, str)
		fmt.Printf("方法1加密后:%s \n", dest1)

		// 解密操作
		dest1 = []byte("xxxxxxxxxxx")
		dest2 := make([]byte, len(dest1))
		cipher2, _ := rc4.NewCipher(key) // 切记：这里不能重用cipher1，必须重新生成新的
		cipher2.XORKeyStream(dest2, dest1)
		fmt.Printf("方法1解密后:%s \n\n", dest2)
	}
}
