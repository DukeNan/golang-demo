/*
学员信息管理系统
1. 添加学生信息
2. 编辑学生信息
3. 展示学生信息
*/

package main

import (
	"fmt"
	"os"
)

// 打印菜单
func showMenu() {
	desc := `
******************************************
          欢迎来到学生信息管理系统             
            1. 添加学生                     
            2. 编辑学生信息
            3. 展示学生信息
            4. 退出系统
******************************************
		`
	fmt.Println(desc)
}

//    获取用户输入的学生信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学生信息")
	fmt.Print("请输入学生学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生班级：")
	fmt.Scanf("%s\n", &class)
	// 调用student的构造函数构造一个学生
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStuMgr()
	for {
		// 打印系统菜单
		showMenu()
		// 接收用户选项
		var input int
		fmt.Print("请输入你要操作的序号：")
		fmt.Scanf("%d\n", &input)
		fmt.Printf("用户的选择是：%d\n", input)
		switch input {
		case 1:
			// 添加学生
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			sm.showStudents()
		case 4:
			// 退出
			os.Exit(0)
		default:
			fmt.Println("您的选择不存在，请重新输入")
		}
	}

}
