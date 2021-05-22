package main

import "fmt"

// 嵌套结构体

// Address结构实体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

// User 用户实体
type User struct {
	Name    string
	Age     int
	Gender  string
	Address Address // 匿名字段
	Email
}

type User1 struct {
	Name    string
	Age     int
	Address // 匿名字段
}

type Email struct {
	Account    string
	CreateTime string
}

func main() {
	user := User{
		Name: "老王",
		Age:  18,
		Address: Address{
			Province: "广东",
			City:     "深圳",
		},
	}
	fmt.Printf("user=%#v\n", user)
	fmt.Println("========嵌套匿名字段=============")
	user1 := User1{
		Name: "老王",
		Age:  18,
		Address: Address{
			Province: "广东",
			City:     "深圳",
		},
	}
	fmt.Printf("user=%#v\n", user1)
	fmt.Println("========嵌套结构体的字段名冲突==============")
	var user3 User
	user3.Name = "张三"
	user3.Gender = "male"
	user3.Email.CreateTime = "1223344444"
	user3.Address.CreateTime = "23234434343"
	fmt.Printf("user=%#v\n", user3)

}
