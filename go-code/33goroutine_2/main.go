package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func write() {
	// lock.Lock()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Microsecond * 10)
	// lock.Unlock()
	rwlock.Unlock()
	wg.Done()
}

func read() {
	// lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Microsecond)
	// lock.Unlock()
	rwlock.RUnlock()
	wg.Done()
}

func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)
	fmt.Println("===========读写互斥锁=============")
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

}
