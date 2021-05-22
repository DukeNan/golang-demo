/*
结构体的“继承”
结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

*/
package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！！！\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 通过匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪叫！！！\n", d.name)
}

type Student struct {
	ID     int8
	Gender string
	Name   string
}

type Class struct {
	Title   string
	Student []*Student
}

func main() {
	d := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "豆豆",
		},
	}
	d.wang()
	d.move()
	fmt.Println("=========结构体与JSON序列化===========")
	c := &Class{
		Title:   "三年二班",
		Student: make([]*Student, 0, 200),
	}
	for i := 1; i <= 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "male",
			ID:     int8(i),
		}
		c.Student = append(c.Student, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Student":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

}
