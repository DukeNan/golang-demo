package main

import "fmt"

func main() {
	// &:取地址
	// *:根据地址取值
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	m := *p
	fmt.Println(m)
}
