package main

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	C    chan string // 用户发送数据的管道
	Name string      // 用户名
	Addr string      // 网络地址
}

var onlineMap map[string]Client

var message = make(chan string)

func Manager() {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message

		for _, cli := range onlineMap {
			cli.c <- msg
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return buf
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	cliAddr := conn.RemoteAddr().String()
	cli := Client(make(chan string), cliAddr, cliAddr)
	onlineMap[cliAddr] = cli
	go WriteMsgToClient(cli, conn)

	//	message <- "[" + cli.Addr + "]" + cli.Name + ": online"
	message <- MakeMsg(cli, "online")
	cli.c <- MakeMsg(cli, "I am here")
	isQuit := make(chan bool)

	go func() {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if n == 0 {
			isQuit <- true
			fmt.Println("conn.Read err = ", err)
			return
		}

		msg := string(buf[:n-1])
		if len(msg) == 3 && msg == "who" {
			conn.Write([]byte("user list:\n"))
			for _, temp := range onlineMap {
				msg = temp.Addr + ":" + temp.Name + "\n"
				conn.Write([]byte(msg))
			}
		} else if len(msg) >= 8 && msg[:6] == "rename" {
			name := strings.Split(msg, "|")[1]
			cli.Name = name
			onlineMap[cliAddr] = cli
			conn.Write([]byte("rename ok\n"))
		} else {
			message <- MakeMsg(cli, msg)
		}

		hasData <- true
	}()

	for {
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, "login out")
			return
		case <-hasData:
		case <-time.After(60 * time.Second):
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, "time out leave out")
			return
		}
	}
}
