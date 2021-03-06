package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	// 函数调用完毕，自动关闭con
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect successful")

	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("[%s] : %s\n", addr, string(buf[:n]))

	if "exit" == string(buf[:n]) {
		fmt.Println("exit")
		return
	}
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		// 处理用户请求，新建一个协程
		go HandleConn(conn)
	}
}
