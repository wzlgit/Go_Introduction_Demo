package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap: a= %d, b= %d\n", a, b)
}

func main() {
	a, b := 10, 20

	// 变量本身传递，值传递
	swap(a, b)
	fmt.Printf("swap: a= %d, b= %d\n", a, b)
	// 输出：swap: a= 20, b= 10
	//  		swap: a= 10, b= 20
}
