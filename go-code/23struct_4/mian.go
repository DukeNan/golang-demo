package main

import "fmt"

// 任意类型添加方法
type MyInt int

// SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int")
}

type Person struct {
	string
	int
}

func main() {
	var m1 MyInt
	m1.SayHello()
	m1 = 100
	fmt.Printf("%#v, %T\n", m1, m1)
	fmt.Println("===========结构体的匿名字段=================")
	p1 := Person{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.string, p1.int)
}
