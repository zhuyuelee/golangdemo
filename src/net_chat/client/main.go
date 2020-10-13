package main

import (
	"fmt"
	"io"
	"net"
	"regexp"
	"strings"
)

var exit chan bool = make(chan bool)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("net.Dial err", err)
	}
	defer conn.Close()

	// 监听消息
	go messageHandler(conn)

	go func() {
		for {
			var input string
			fmt.Scanln(&input)
			if ok, _ := regexp.MatchString(`^exit$`, input); ok {
				exit <- true
				break
			}
			if strings.Index(input, "::") > -1 {
				msgs := strings.Split(input, "::")
				fmt.Printf("您对[%s]说：%s\n", msgs[0], msgs[1])
			} else {
				fmt.Println("您说：", input)
			}
			//发送消息
			conn.Write([]byte(input))
		}
	}()

	<-exit
}

func messageHandler(conn net.Conn) {
	for {
		buf := make([]byte, 1024*2)
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("conn.Read err", err)
		}
		if n == 0 {
			fmt.Println("服务器断开了")
			exit <- true
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
