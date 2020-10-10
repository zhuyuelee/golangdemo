package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//checkInputEmail()
	regNum()
}

//提取用户多输入多个数字
func regNum() {
	var reg = regexp.MustCompile(`(\d+(.\d+)?)(;\d+(.\d+)?)*`)
	for {
		var input string
		fmt.Print("请输入数字(多个数据请;分隔。E退出)：")
		fmt.Scanln(&input)
		if ok, _ := regexp.MatchString("^[Ee]$", input); ok {
			break
		}
		if nums := reg.FindAllString(input, -1); len(nums) > 0 {
			fmt.Println("输入的数字有：" + strings.Join(nums, ","))
		} else {
			fmt.Println("数字输入错误")
		}
	}
}

func checkInputEmail() {
	var reg = regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	for {
		var email string
		fmt.Print("请输入邮箱(E退出)：")
		fmt.Scanln(&email)
		if ok, _ := regexp.MatchString("^[Ee]$", email); ok {
			break
		}
		if ok := reg.MatchString(email); ok {
			fmt.Println("邮箱输入正确：" + email)
		} else {
			fmt.Println("邮箱格式错误")
		}
	}

}
