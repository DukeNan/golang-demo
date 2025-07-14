package main

import (
	"fmt"
	// "runtime"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutin!", i)
}

func a() {
	for i := 1; i <= 10; i++ {
		fmt.Println("A:", i)
	}
	wg.Done()
}

func b() {
	for i := 1; i <= 10; i++ {
		fmt.Println("B:", i)
	}
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()

	// 设置cpu的核数
	// runtime.GOMAXPROCS(8)
	// wg.Add(2)
	// go a()
	// go b()
	// wg.Wait()

}
