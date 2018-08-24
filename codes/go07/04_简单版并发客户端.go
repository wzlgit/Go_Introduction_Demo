package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	defer conn.Close()

	// 接收服务器回复的数据
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read err = ", err)
				return
			}
			fmt.Println(buf[:n])
		}
	}()

	for {
		str := make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("conn.Stdin.Read err = ", err)
			return
		}

		// 把输入的内容发送给服务器
		conn.Write(str[:n])
	}
}
