// 求素数
// 一个管道里放数字
// 一个管道里放素数
// 开4个协程，共同操作

package main

import "fmt"

func putNumber(numChan chan int) {
	for i := 1; i <= 20; i++ {
		numChan <- i
	}
	// 关闭管道
	close(numChan)
}

func isPrime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func getPrime(numChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range numChan {
		if isPrime(num) {
			primeChan <- num
		}
	}
	exitChan <- true
}

func multiGoroutineGetPrime(numChan chan int, primeChan chan int, exitChan chan bool) {

	for i := 0; i < 4; i++ {
		go getPrime(numChan, primeChan, exitChan)
	}
}

func closePrime(primeChan chan int, exitChan chan bool) {
	for i := 0; i < 4; i++ {
		<-exitChan
	}
	close(primeChan)
}

func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println("素数是", v)
	}
}
