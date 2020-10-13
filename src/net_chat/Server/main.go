package main

import (
	"fmt"
	"io"
	"net"
	"regexp"
	"strings"
)

type message struct {
	msg, from, to string
}

var clientConn map[string]net.Conn
var inputChat chan message

func init() {
	inputChat = make(chan message, 10)
	clientConn = make(map[string]net.Conn, 10)
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(" net.Listen error:", err)
	}
	defer listen.Close()

	defer func() {
		for _, conn := range clientConn {
			conn.Close()
		}
	}()
	fmt.Println("服务启动成功，等待接入")

	//发送消息
	go sendMsg()
	for {
		conn, connErr := listen.Accept()
		if connErr != nil {
			fmt.Println(" listen.Accept error:", connErr)
		}
		clientJoin(conn)
	}

}

//clientJoin 欢迎加入
func clientJoin(conn net.Conn) {
	clienKey := conn.RemoteAddr().String()
	fmt.Println(clienKey, "接入")
	inputChat <- message{fmt.Sprintf("%s 加入聊天", clienKey), clienKey, ""}
	clientConn[clienKey] = conn
	inputChat <- message{fmt.Sprintf("weclome %s", clienKey), "", clienKey}

	go receiveMsg(conn)
}

// sendMsg 给客户发消息
func sendMsg() {
	for {
		msg := <-inputChat
		fmt.Printf("发送消息:%s \n", msg)
		//给一个人发消息
		if msg.to != "" {
			if conn, ok := clientConn[msg.to]; ok {
				if msg.from != "" {
					conn.Write([]byte(fmt.Sprintf("[%s]对您说：%s", msg.from, msg.msg)))
				} else {
					conn.Write([]byte(msg.msg))
				}
			}
		} else {
			//给所有人发消息
			for key, conn := range clientConn {
				if msg.from != key {
					conn.Write([]byte(fmt.Sprintf("[%s]%s", msg.from, msg.msg)))
				}
			}
		}
	}
}

//receiveMsg 监听用户数据
func receiveMsg(conn net.Conn) {
	for {
		buf := make([]byte, 1024*2)
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("conn.Read err", err)
		}
		var clientKey = conn.RemoteAddr().String()
		if n == 0 { //对方断开，或者，出问题
			delete(clientConn, clientKey)
			inputChat <- message{fmt.Sprintf("%s 离开了", clientKey), clientKey, ""}
			return
		}

		input := string(buf[:n])
		fmt.Printf("收到%s消息：%s\n", clientKey, input)
		//特殊命令
		if ok, _ := regexp.MatchString(`^who$`, input); ok {
			var keys string
			for key := range clientConn {
				keys += key + "\n"
			}
			inputChat <- message{strings.TrimRight(keys, "\n"), "", clientKey}
		} else {
			msg, to := "", ""
			if strings.Index(input, "::") > -1 {
				msgs := strings.Split(input, "::")
				if len(msgs) == 2 {
					msg, to = msgs[1], msgs[0]
				} else {
					msg = input
				}
			} else {
				msg = input
			}

			inputChat <- message{msg, clientKey, to}
		}
	}

}
