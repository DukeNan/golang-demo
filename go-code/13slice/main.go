package main

import (
	"fmt"
)

func main() {
	// 定义
	// var a []string
	// var b = []int{}
	// var c = []bool{true, false}
	// var d = []bool{true, false}
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(a == nil)
	// fmt.Println(b == nil)

	fmt.Println("==============切片表达式================")
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", s)
	fmt.Printf("s=%v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s))
	fmt.Printf("%v\n", a[2:])
	fmt.Printf("%v\n", a[:3])
	fmt.Printf("%v\n", a[:])
	fmt.Printf("%v\n", a[:])
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Println(s2)
	println("============完整切片表达式============================")
	// 完整切片表达式
	b := [5]int{1, 2, 3, 4, 5}
	t := b[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	fmt.Println("========使用make()函数构造切片============= ")
	m := make([]int, 2, 10)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	fmt.Println(len(m) == 0)

	fmt.Println("==========切片的赋值拷贝=============")
	m1 := make([]int, 3)
	m2 := m1
	m2[0] = 100
	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println("======切片遍历===============")
	n := []int{1, 3, 5}
	fmt.Println(n, len(n), cap(n))
	for i := 0; i < len(n); i++ {
		fmt.Println(i, n[i])
	}
	for _, v := range n {
		fmt.Println(v)
	}

	fmt.Println("=========append============")
	var p []int
	p = append(p, 1)       // [1]
	p = append(p, 2, 3, 4) // [1 2 3 4]
	q2 := []int{5, 6, 7}
	p = append(p, q2...) // [1 2 3 4 5 6 7]
	fmt.Println(p)
	fmt.Println(q2)

	// 通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	var q []int
	q = append(q, 9, 8, 7)
	fmt.Println(q)
	fmt.Println(len(q), cap(q))
	q = append(q, 6, 5, 4)
	fmt.Println(len(q), cap(q))

	// 每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
	// 当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
	// “扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
	fmt.Println("=========append()添加元素和切片扩容=================")
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	var cityS []string
	cityS = append(cityS, "深圳")
	cityS = append(cityS, "北京", "上海", "广州")
	cs := []string{"成都", " 武汉"}
	cityS = append(cityS, cs...)
	fmt.Println(cityS)

	fmt.Println("========copy============")
	u := []int{1, 2, 3, 4, 5}
	v := make([]int, 5, 6)
	copy(v, u)
	fmt.Println(u, v)
	v[0] = 1000
	fmt.Println(u, v)

	fmt.Println("========删除元素=============")
	x := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	x = append(x[:2], x[3:]...)
	fmt.Println(x)
}
