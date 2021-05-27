package main

import "fmt"

type Sayer interface {
	say()
}

type dog struct{}

type cat struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")

}

type Move interface {
	move()
}

func (d dog) move() {
	fmt.Println("狗会动")
}

type Bark interface {
	bark()
}

func (d *dog) bark() {
	fmt.Println("狗会叫")
}

func main() {
	// 接口类型变量能够存储所有实现了该接口的实例
	var s Sayer // 声明一个Sayer类型的变量x
	c := cat{}  // 实例化一个cat
	d := dog{}  // 实例化一个dog
	s = c       // 可以把cat实例直接赋值给s
	s.say()     // 喵喵喵
	s = d       // 可以把dog实例直接赋值给s
	s.say()     // 汪汪汪
	fmt.Println("=======值接收者和指针接收者=============")
	var m Move
	wangcai := dog{}
	m = wangcai
	fuigui := &dog{}
	m = fuigui
	m.move()
	fmt.Println("========")
	var b Bark
	var doudou = dog{}
	b = &doudou
	b.bark()

}
