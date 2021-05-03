package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func f10() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 1.返回值=x, 执行defer， 返回5
}

func f20() (x int) { // 返回x
	defer func() {
		x++
	}()
	return 5 // 1.返回值=x=5, 2.执行defer，x=6, 返回x=6
}

func f30() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 1.返回值=y=5, 2.执行defer x=6, 3.返回y=5
}

func f40() (x int) {
	defer func(x int) {
		x++ // 改变的是函数的副本
	}(x) // x作为参数传递给函数
	return 5 //1. 返回值=x=5, 2.执行defer，3.返回x=5
}

func calculate(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量作为立即执行函数
func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20)
	// 自执行函数：匿名函数定义完家()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
	fmt.Println("========闭包==============")
	var f = adder()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(30))

	f1 := adder()
	fmt.Println(f1(40))
	fmt.Println(f1(50))
	fmt.Println("==========示例2============")
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc(".jpg"))
	fmt.Println(txtFunc("test"))
	fmt.Println("========示例3============")
	m, n := calc(10)
	fmt.Println(m(1), n(2))
	fmt.Println("========defer==============")
	// 分三部：
	// 1: 返回值赋值
	// 2: 执行defer
	// 3: 返回真正的返回值
	// fmt.Println("statr")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("end")
	fmt.Println("==========defer测试==========")
	fmt.Println(f10())
	fmt.Println(f20())
	fmt.Println(f30())
	fmt.Println(f40())
	fmt.Println("=======defer面试==============")
	// defer
	x := 1
	y := 2
	defer calculate("AA", x, calculate("A", x, y))
	x = 10
	defer calculate("BB", x, calculate("B", x, y))
	y = 20

}
