package main

import "fmt"

var a int

func main() {
	a = 10
	fmt.Println("a = ", a)

	test()
}

func test() {
	fmt.Println("test a = ", a)
}
