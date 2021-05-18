package main

import (
	"fmt"
	"unsafe"
)

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

// 定义结构体
type person struct {
	name string
	city string
	age  int8
}

type person1 struct {
	name, city string
	age        int8
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)
	fmt.Println("==========实例化=============")

	// 实例化结构体
	var p1 person
	p1.name = "laowang"
	p1.city = "深圳"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	fmt.Println("========匿名结构体=============")
	var user struct {
		name string
		age  int8
	}
	user.name = "老王"
	user.age = 20
	fmt.Printf("user=%v\n", user)

	fmt.Println("========指针类型结构体=============")
	// 使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(person)
	fmt.Printf("%T\n", p2) // p2是一个指针类型
	fmt.Printf("p2=%#v\n", p2)
	// 对结构体指针直接使用.来访问结构体的成员
	fmt.Println(p2.age) // 默认为go的零值
	p2.name = "shaun"
	p2.age = 20
	p2.city = "shenzhen"
	fmt.Printf("name=%s, age=%d, city=%s\n", p2.name, p2.age, p2.city)

	fmt.Println("=======取结构体地址实例化==============")
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "七米"
	p3.age = 30
	p3.city = "Guangzhou"
	fmt.Printf("p3=%#v\n", p3) // p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。

	fmt.Println("==========结构体初始化====================")
	// 使用键值对初始化
	p5 := person{
		name: "王五",
		city: "北京",
		age:  30,
	}
	fmt.Printf("p5=%#v\n", p5)
	// 对结构体指针进行键值对初始化,  字段可以省略，默认零值
	p6 := &person{
		name: "小王子",
		age:  20,
	}
	fmt.Printf("p6=%#v\n", p6)
	// 使用值列表初始化
	p8 := &person{
		"古力娜扎",
		"北京",
		30,
	}
	fmt.Printf("p8=%#v\n", p8)

	fmt.Println("==========结构体内存布局===============")
	n := test{1, 2, 3, 4} //结构体占用一块连续的内存
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	fmt.Println("========空结构体==============")
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 空结构体不占内存
}
