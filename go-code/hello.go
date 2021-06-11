package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("程序结束...")
}
