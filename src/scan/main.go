package main

import "fmt"

func main() {
	var input string
	fmt.Print("请输入：")
	//scan输入要传入指针，否则跳过输入
	fmt.Scanln(&input)
	fmt.Printf("输入了:%s", input)
}
