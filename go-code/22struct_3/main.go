package main

import "fmt"

//Person 结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Deram Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Golang\n", p.name)
}

// 指针类型接收者
func (p *Person) SetAge(newAge int8)  {
	p.age = newAge
}

// 值类型接收者
func (p Person) SetAge2(newAge int8){
	p.age = newAge
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
	fmt.Printf("%s的年龄是：%d\n", p1.name, p1.age)
	p1.SetAge(30)
	fmt.Printf("%s的年龄是：%d\n", p1.name, p1.age)
	p1.SetAge2(35)
	fmt.Printf("%s的年龄是：%d\n", p1.name, p1.age)
}
