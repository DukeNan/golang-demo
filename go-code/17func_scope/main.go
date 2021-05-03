package main

import (
	"errors"
	"fmt"
)

// 定义全局变量
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num)
}

func testLocalvar() {
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

func testNum() {
	// 有限访问局部变量
	num := 100
	fmt.Printf("num=%d\n", num)
}

func testLocalvar2(x, y int) {
	fmt.Println(x, y)
	if x > 10 {
		z := 100 // 只在if语句块中生效
		fmt.Println(z)
	}
}

func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) // 变量i只在当前for语句块中生效
	}
}

func add(x, y int) int {
	sum := x + y
	return sum
}

func sub(x, y int) int {
	return x - y
}

type calculation func(int, int) int

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别字符串")
		return nil, err
	}
}

func main() {
	testGlobalVar()
	testLocalvar()
	testNum()
	testLocalvar2(10, 20)
	testLocalVar3()
	fmt.Println("========函数类型和变量============")
	var c calculation = add
	ret := c(10, 20)
	fmt.Printf("%T\n", c)
	fmt.Println(ret)
	f := add
	fmt.Printf("type of f:%T\n", f)
	fmt.Println(f(5, 10))
	// 函数作为参数
	fmt.Println("=========函数作为参数==========")
	ret2 := calc(2, 3, add)
	fmt.Println(ret2)
	fmt.Println("=======函数作为返回值============")
	ret3, _ := do("+")
	fmt.Println(ret3(10, 100))

}
