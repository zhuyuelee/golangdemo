package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出错了", err)
		}
	}()

	f1(10)
	f2()
}

func f1(index int) {
	s := [3]int{1, 2, 3}
	s[index] = 100
}
func f2() {
	fmt.Println("我是f2")
}
