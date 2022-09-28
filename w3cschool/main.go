package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()

	rwlock.RLock()
	fmt.Println("Read x= ", x)
	time.Sleep(time.Millisecond)

	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	fmt.Println("Write x=", x)
	x += 1
	time.Sleep(time.Millisecond)
	fmt.Println("Write x+1=", x)
	rwlock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 20; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("----------------")
	fmt.Println(time.Since(start))
}
