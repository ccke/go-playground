package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 50000,
	})
	if err != nil {
		fmt.Println("消息解码失败, 错误信息:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("读取UDP数据失败, 错误信息：", err)
			continue
		}
		fmt.Printf("数据：%v 地址：%v 数据大小：%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("写UDP数据失败, 错误信息：", err)
			continue
		}
	}
}
