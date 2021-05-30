package main

import (
	"fmt"
	"github.com/ccke/go-playground/study_03/play_01/tcp_socket/proto"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000") // 发起连接
	if err != nil {
		fmt.Println("建立连接失败，错误信息：", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 5; i++ {
		msg := "你好，在吗？"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("消息编码失败, 错误信息:", err)
			return
		}
		conn.Write(data)
	}
}

