package student

import "fmt"

// Student 学生
type Student struct {
	Name      string
	Age       int
	ClassRoom string
}

var studentList = make(map[string]Student, 50)

//InitData 初始化数据
func InitData() {
	studentList["1001"] = Student{Name: "李明", Age: 10, ClassRoom: "1.1班"}
	studentList["1002"] = Student{Name: "赵四", Age: 11, ClassRoom: "1.3班"}
	studentList["1003"] = Student{Name: "王六", Age: 12, ClassRoom: "1.4班"}
}

// GetStudent 获取学生列表
func GetStudent() {
	fmt.Println("学号	姓名	年龄	班级")
	for key, stu := range studentList {
		fmt.Printf("%s	%s	%d岁	%s\n", key, stu.Name, stu.Age, stu.ClassRoom)
	}
}

//Add 添加学生
func (s Student) Add(no string) (result bool) {
	if _, ok := studentList[no]; ok {
		fmt.Printf("学号为%s的学生已存在\n", no)
		result = false
	} else {
		studentList[no] = s
		result = true
	}
	return
}

//Edit 修改学生
func (s Student) Edit(no string) (result bool) {
	if _, ok := studentList[no]; ok {
		studentList[no] = s
		result = true
	} else {
		fmt.Printf("学号为%s的学生不存在\n", no)
		result = false
	}
	return
}

//Delete 删除学生
func Delete(no string) (result bool) {
	if _, ok := studentList[no]; ok {
		delete(studentList, no)
		result = true
	} else {
		fmt.Printf("学号为%s的学生不存在\n", no)
		result = false

	}
	return
}
