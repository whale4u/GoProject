package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

//rsa加密解密 签名验签
func main() {
	//生成私钥
	priv, e := rsa.GenerateKey(rand.Reader, 1024)
	if e != nil {
		fmt.Println(e)
	}

	//根据私钥产生公钥
	pub := &priv.PublicKey
	////也可以这样写
	//pub := priv.PublicKey

	//明文
	//plaintext := []byte("Hello world")
	file, _ := os.Open("./1.txt")
	//file, _ := os.Open("./1.png")
	defer file.Close()

	plaintext, _ := ioutil.ReadAll(file)

	//加密生成密文
	fmt.Printf("%q\n加密:\n", plaintext)
	ciphertext, e := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, plaintext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%x\n", ciphertext)

	//解密得到明文
	fmt.Printf("解密:\n")
	plaintext, e = rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%q\n", plaintext)

	//写入文件
	err := ioutil.WriteFile("./2.txt", plaintext, 0644)
	//err := ioutil.WriteFile("./2.png", plaintext, 0644)
	if e != nil {
		panic(err)
	}

	//消息先进行Hash处理
	h := sha256.New()
	h.Write(plaintext)
	hashed := h.Sum(nil)
	fmt.Printf("%q sha256 Hashed:\n\t%x\n", plaintext, hashed)

	//签名
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.SHA256}
	sig, e := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, hashed, opts)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("签名:\n\t%x\n", sig)

	//认证
	fmt.Printf("验证结果:")
	if e := rsa.VerifyPSS(pub, crypto.SHA256, hashed, sig, opts); e != nil {
		fmt.Println("失败:", e)
	} else {
		fmt.Println("成功.")
	}
}
