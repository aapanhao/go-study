// 实践协程求素数
// 使用管道，共享空间（替代加锁）
// 使用select 替代close一个管道

package main

func testDemo() {
	// 放数字
	// 空间可以开的比较小，编译器判断只要有协程从里取值，就不会报panic（超出空间大小）
	// 不过要使用协程来放数
	numChan := make(chan int, 10)
	go putNumber(numChan)

	//素数chan
	primeChan := make(chan int, 20)
	//需要有一个close的逻辑，要不然会报panic
	exitChan := make(chan bool, 4)
	go multiGoroutineGetPrime(numChan, primeChan, exitChan)

	//输出
	//需要有一个close的逻辑，要不然会报panic
	closePrime(primeChan, exitChan)
	printPrime(primeChan)
}

func testLock() {
	multiAdd()
}

func testManual() {
	testManualClose()
}

func main() {
	testSafeSend()
}
