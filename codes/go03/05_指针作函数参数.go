package main

import "fmt"

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Printf("swap: *p1= %d, *p2= %d\n", *p1, *p2)
}

func main() {
	a, b := 10, 20

	// 通过一个函数交换a和b的内容
	swap(&a, &b)
	fmt.Printf("swap: a= %d, b= %d\n", a, b)
	// 输出：swap: *p1= 20, *p2= 10
	//		swap: a= 20, b= 10
}
