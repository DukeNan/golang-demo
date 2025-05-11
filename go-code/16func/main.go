package main

import "fmt"

// 函数定义
func getName(age int) string {
	return "laowang" + fmt.Sprintf("%d", age)
}

// 函数定义
func intSum(x, y int) int {
	return x + y
}

// 可变参数
func intSum2(x ...int) int {
	fmt.Println(x) // x是切片
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 固定参数搭配可变参数
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

// 多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 多返回值	
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 当我们的一个函数返回值类型为slice时，nil可以看成一个有效的slice，没有必要显示返回一个长度为0的切片。
func someFunc(x string) []int {
	if x == "" {
		return nil
	}
	return []int{1, 2, 3, 4, 5, 6}
}

func main() {
	// 函数调用
	name := getName(20)
	fmt.Println(name)
	total := intSum(10, 20)
	fmt.Println(total)
	fmt.Println("==========可变参数============")
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4)
	fmt.Println("=========固定参数搭配可变参数=================")
	ret5 := intSum3(100)
	ret6 := intSum3(100, 10)
	ret7 := intSum3(100, 10, 20)
	ret8 := intSum3(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8)
	fmt.Println("=======多返回值================")
	a, b := calc(10, 2)
	fmt.Println(a, b)
	fmt.Println("=======返回值命名================")
	c, d := calc1(7, 3)
	fmt.Println(c, d)
	// 当我们的一个函数返回值类型为slice时，nil可以看成一个有效的slice，没有必要显示返回一个长度为0的切片。
	ret9 := someFunc("")
	fmt.Println(ret9)

}
