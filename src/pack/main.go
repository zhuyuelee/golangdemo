package main

import (
	"fmt"
	"pack/student"
)

func main() {
	student.InitData()
	for {
		fmt.Println("-----------------学生管理-----------------")
		student.GetStudent()
		fmt.Println("-----------------------------------------")
		fmt.Println(`-----------------管理菜单-----------------
1、添加学生
2、修改学生
3、删除学生
		`)
		fmt.Println("请输入(0退出)：")
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			addStudent()
		case 2:
			editStudent()
		case 3:
			deleteStudent()
		case 0:
			break
		default:
			fmt.Println("输入错误")
		}
	}
}

func inputStudentNo() (no string) {
	print("请输入学生学号：")
	fmt.Scanln(&no)
	return
}

func inputStudent() student.Student {
	var (
		name      string
		classRoom string
		age       int
	)
	print("请输入学生姓名：")
	fmt.Scanln(&name)
	print("请输入学生年龄：")
	fmt.Scanln(&age)
	print("请输入学生班级：")
	fmt.Scanln(&classRoom)

	return student.Student{
		Name:      name,
		Age:       age,
		ClassRoom: classRoom,
	}
}

func addStudent() {
	no := inputStudentNo()
	student := inputStudent()
	student.Add(no)
}

func editStudent() {
	no := inputStudentNo()
	student := inputStudent()
	student.Edit(no)
}

func deleteStudent() {
	no := inputStudentNo()
	student.Delete(no)
}
