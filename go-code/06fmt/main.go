package main

import "fmt"

// 占位符
func main() {
	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

	var f = 1.23855
	fmt.Printf("%T\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%#v\n", f)
}
