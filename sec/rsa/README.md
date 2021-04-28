## 简介
首先不得不提的一点，RSA是一个确定性的算法，也就是说，给定密钥，同样的明文总是会计算出同样的密文。
假如Jim知道明文大概会是什么样的，比如11位的电话号码，他就可以用少得多的代价，尝试出所有的明文-密文组合，
也就可以在下一次拿到密文的时候，对比出相应的明文。

为了保证择密文攻击下（chosen cipher-text attack）的安全性，我们首先要将算法改造为非确定性，
具体的做法就是通过“Padding”，就是将明文与某个随机数组合起来，这样明文就被转移到了一个更大的明文空间上，一定程度上防止了选择明文攻击。
但这样还是不够，两千年前后密码学界提出了更安全的Padding模式，即RSA-PSS和PSA-OAEP，
后者正是PKCS#1 v2中使用RSA的标准方式，详见：

## 最优非对称加密填充OAEP
OAEP（Optimal Asymmetric Encryption Padding）
```go
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)
func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
```

## 关注点
RSA只能加密小于(或等于) key 长度的数据。
如果要加密大文件，请使用使用对称算法(例如AES)加密数据。
如果需要RSA公钥/私钥对，请使用RSA加密对称算法(例如AES)的密钥文件 —— 这称为混合加密，本质上是HTTPS如何加密数据。

参考文章：
- https://golang.org/pkg/crypto/rsa/#DecryptOAEP
- https://zh.wikipedia.org/wiki/%E6%9C%80%E4%BC%98%E9%9D%9E%E5%AF%B9%E7%A7%B0%E5%8A%A0%E5%AF%86%E5%A1%AB%E5%85%85