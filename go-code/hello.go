package main

import "fmt"

func main() {
	// fmt.Println("欢迎来到学生信息管理系统")
	// fmt.Println("1. 添加学生")
	// fmt.Println("2. 编辑学生信息")
	// fmt.Println("3. 展示学生信息")
	// fmt.Println("4. 退出系统")
	// fmt.Println("\n
	// *******************
	// ")
	desc := `
*******************************************
          欢迎来到学生信息管理系统             
             1. 添加学生                     
             2. 编辑学生信息
             3. 展示学生信息
             4. 退出系统
*******************************************
	`
	fmt.Println(desc)
	fmt.Printf("%T\n", desc)
}
