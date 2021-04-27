package main

import "fmt"

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 十进制 --> 二进制
	fmt.Printf("%o\n", i1) // 十进制 --> 八进制
	fmt.Printf("%x\n", i1) // 十进制  --> 十六进制
	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	// 查看变量类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型
	i4 := int8(9)   // 明确指定int8类型，否则默认int类型
	fmt.Printf("%T\n", i4)
}





