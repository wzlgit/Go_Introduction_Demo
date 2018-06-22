package main

import "fmt"

func testa() {
	fmt.Println("aaa")
}

func testb() {
	panic("this a panic test")
	fmt.Println("bbb")
}

func testc() {
	fmt.Println("ccc")
}

func main() {
	testa()
	testb()
	testc()
}
