package main

import "fmt"

var a byte // 全局变量

func main() {
	//	var a int // 局部变量	// 注释掉则输出：uint8

	// 1、不同作用域，允许定义同名变量
	// 2、使用变量的原则：就近原则
	fmt.Printf("%T\n", a) // 输出：init

	{
		var a float32
		fmt.Printf("2.%T\n", a) // 输出：float32
	}
}
