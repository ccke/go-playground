package main

import (
	"bufio"
	"fmt"
	"github.com/ccke/go-playground/study_03/play_01/tcp_socket/proto"
	"io"
	"net"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("消息解码失败, 错误信息:", err)
			break
		}
		fmt.Println("收到客户端发来的数据：", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000") // 监听
	if err != nil {
		fmt.Println("监听失败，错误信息：", err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("建立TCP连接失败，错误信息：", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
