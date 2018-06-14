package main

import "fmt"

func main() {
	var p *int
	p = new(int)
	*p = 666
	fmt.Println("*p = ", *p)
	// 输出：*p =  666

	q := new(int) // 自动推导类型
	*q = 777
	fmt.Println("*q = ", *q)
	// 输出：*q =  777
}
