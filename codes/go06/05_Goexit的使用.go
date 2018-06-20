package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccc")

	// 终止协程
	runtime.Goexit()

	fmt.Println("ddd")
}

func main() {
	// 创建新建的协程
	go func() {
		fmt.Println("aaa")

		test()

		fmt.Println("bbb")
	}()

	// 特地写一个死循环，目的不让主协程结束
	for {

	}
	// 注释runtime.Goexit()时输出：
	// aaa
	// ddd
	// ccc
	// bbb

	// 使用runtime.Goexit()时输出：
	// aaa
	// ccc
}
