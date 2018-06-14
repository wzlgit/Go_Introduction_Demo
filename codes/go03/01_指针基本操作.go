package main // 必须有个main包

import "fmt"

func main() {
	var a int = 10
	// 每个变量有两层含义：变量的内存，变量的地址
	// 运行Go快捷键：Windows->Control+R，Mac->Command+R
	fmt.Printf("a = %d\n", a)
	fmt.Printf("&a = %v\n", &a)

	// 保存某个变量的地址，需要指针类型
	// *int 保存int的地址， **init 保存 *int的地址
	// 声明（定义），定义只是特殊的声明
	// 定义一个变量p，类型为*int
	var p *int
	// 指针变量指向谁，就是把谁的地址赋值给指针变量
	p = &a

	fmt.Printf("p = %v, & a = %v\n", p, &a)
	*p = 666 // *p操作的不是p的内存，而是p所指向的内存（即a）
	fmt.Printf("*p = %v, a = %v\n", *p, a)

	/*
		 * 输出：
		    *   a = 10
			*	&a = 0xc420014100
			*	p = 0xc420014100, & a = 0xc420014100
			*	*p = 666, a = 666
		 *
	*/
}
