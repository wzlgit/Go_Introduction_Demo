package main

import (
	"fmt"
	"time"
)

// 全局变量，创建一个channel
var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

// person1执行完后，才轮到person2执行
func person1() {
	Printer("hello")
	ch <- 666 // 给管道写数据，发送
}

func person2() {
	// 从管道取数据，接收，如果通道没有数据就会阻塞
	<-ch
	Printer("world")
}

func main() {
	//	Printer("hello")
	//	Printer("world")

	//	// hello
	//	// world

	go person1()
	go person2()

	// 特地不让主协程结束，死循环
	for {

	}
}
