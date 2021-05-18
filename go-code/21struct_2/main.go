package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu // 将临时变量stu的地址给map，最后所有的地址都为同样的一个值
		// m[stu.name] = &stu   // 正确做法
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
