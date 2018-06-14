package main

import "fmt"

// p指向实现数组a，它指向数组，是数组指针
// *p代表指针指向的内存，就是实参a
func modify(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("main: *p =", *p)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// 地址传递
	modify(&a)
	fmt.Println("main: a =", a)
	// 输出：
	// main: *p = [666 2 3 4 5]
	// main: a = [666 2 3 4 5]
}
