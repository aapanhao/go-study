package main

import (
	"fmt"
	"sync"
	"time"
)

func fib(mychan chan int) {
	a, b := 1, 1
	n := cap(mychan)
	for i := 0; i < n; i++ {
		mychan <- b
		a, b = b, a+b
	}
	close(mychan)
}

func testFib() {
	mychan := make(chan int, 10)
	go fib(mychan)

	for n := range mychan {
		fmt.Println("number: ", n)
	}
}

func testLock() {

	lockChan := make(chan bool, 1)

	x := 0

	for i := 0; i < 1000; i++ {
		go incr(lockChan, &x)
	}

	time.Sleep(time.Second)

	fmt.Println(x)
}

func incr(lockChan chan bool, x *int) {
	//lockChan <- true
	*x = *x + 1
	//<-lockChan
}

func add(x *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*x = *x + 1
		lock.Unlock()
	}
	wg.Done()
}

func multiAdd() {
	var wg sync.WaitGroup
	lock := &sync.Mutex{}

	wg.Add(3)
	x := 0
	go add(&x, &wg, lock)
	go add(&x, &wg, lock)
	go add(&x, &wg, lock)

	wg.Wait()

	fmt.Print("number", x)

}

func main() {

	pipeline := make(chan int)

	go func() {
		fmt.Println("im", <-pipeline)
	}()

	pipeline <- 1
	close(pipeline)
	time.Sleep(time.Second)
	fmt.Println("last", <-pipeline)

}
