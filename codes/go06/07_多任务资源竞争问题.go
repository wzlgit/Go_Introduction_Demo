package main

import (
	"fmt"
	"time"
)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func person1() {
	Printer("hello")
}

func person2() {
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
