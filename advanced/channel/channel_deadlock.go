package main

//使用信道需要考虑的一个重点是死锁。
// 当 Go 协程给一个信道发送数据时，照理说会有其他 Go 协程来接收数据。
// 如果没有的话，程序就会在运行时触发 panic，形成死锁。
func main() {
	ch := make(chan int)
	ch <- 5
}
