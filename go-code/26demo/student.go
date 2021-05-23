package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

// newStudent是 student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学生管理的类型
type stuMgr struct {
	allStudents []*student
}

// newStuMgr 是stuMgr的构造函数
func newStuMgr() *stuMgr {
	return &stuMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 添加学生
func (s *stuMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 编辑学生
func (s *stuMgr) modifyStudent(stu *student) {
	for i, v := range s.allStudents {
		if stu.id == v.id {
			s.allStudents[i] = stu
			return
		}
	}
	fmt.Printf("学号为：%d的学生不存在", stu.id)
}

// 展示学生
func (s *stuMgr) showStudents() {
	for _, stu := range s.allStudents {
		fmt.Printf("学号：%d， 姓名：%s，班级：%s\n", stu.id, stu.name, stu.class)
	}
}
