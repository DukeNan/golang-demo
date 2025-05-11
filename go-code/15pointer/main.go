package main

import "fmt"

func main() {
	// &:取地址
	// *:根据地址取值
	n := 18
	p := &n
	// 打印p的值
	fmt.Println(p)
	// 打印p的类型
	fmt.Printf("%T\n", p)
	// 将p的值赋值给m
	m := *p
	// 打印m的值
	fmt.Println(m)
	
	
}
