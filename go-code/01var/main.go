package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "laowang"
	age = 18
	isOk = true
	// Go语言中变量声明必须使用， 不使用编译不过去
	fmt.Print(isOk)              // 在终端输出要打印的内容
	fmt.Println()
	fmt.Printf("name:%s\n", name) //  格式化输出
	fmt.Println(age)             // 打印之后换行
	fmt.Println("hello world")

	//  声明变量同时赋值
	var s1 string = "lisi"
	fmt.Println(s1)

	// 类型推导（根据值判断该变量的类型）
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明,  只能在函数里面声明
	s3 := "hahaha"
	fmt.Println(s3)

	// S1 := "10"  //   同一个作用域不能重复声明同名变量

	//  匿名变量是一个特殊的变量：_
}
