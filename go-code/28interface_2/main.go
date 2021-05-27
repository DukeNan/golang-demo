package main

import "fmt"

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer // 嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) move() {
	fmt.Printf("%s会动...\n", c.name)
}

func (c cat) say() {
	fmt.Printf("%s喵喵喵...\n", c.name)
}

func main() {
	var w WashingMachine
	h := haier{}
	w = h
	w.wash()
	w.dry()
	fmt.Println("======接口嵌套============")
	var a animal = cat{name: " 小花"}
	a.move()
	a.say()
	fmt.Println("=======空接口===========")
	var x interface{}
	s := "Hello World"
	x = s
	fmt.Printf("type: %T, value: %v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type: %T, value: %v\n", x, x)
	b := true
	x = b
	fmt.Printf("type: %T, value: %v\n", x, x)

}
