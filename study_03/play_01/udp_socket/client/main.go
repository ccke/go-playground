package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 50000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，错误信息：", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，错误信息：", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，错误信息：", err)
		return
	}
	fmt.Printf("数据：%v 地址：%v 数据大小：%v\n", string(data[:n]), remoteAddr, n)
}

