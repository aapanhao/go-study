package main

import (
	"fmt"
	"sync"
)

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

	fmt.Println("number", x)
}


