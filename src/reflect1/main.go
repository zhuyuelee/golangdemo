package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
}

func main() {
	stu := student{"张三"}

	t := reflect.TypeOf(stu)
	v := reflect.ValueOf(stu)
	fmt.Println(t)
	//转换为原来的类型
	if value, ok := v.Interface().(student); ok {
		fmt.Println(value.Name)
	}
}
