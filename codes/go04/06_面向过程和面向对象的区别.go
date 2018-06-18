package main

import "fmt"

// 面向过程
func Add01(a, b int) int {
	return a + b
}

// 面向对象
type long int

func (temp long) Add02(other long) long {
	return temp + other
}

func main() {
	var result int
	result = Add01(1, 2)
	fmt.Println("result = ", result)
	// result =  3

	var a long = 2
	result2 := a.Add02(3)
	fmt.Println("result2 =", result2)
	// result2 = 5
}
