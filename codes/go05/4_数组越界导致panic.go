package main

import "fmt"

func testa() {
	fmt.Println("aaa")
}

func testb(x int) {
	var a [10]int
	a[x] = 111
	// 当x为20时，导致数组越界，产生一个panic，导致程序崩溃
}

func testc() {
	fmt.Println("ccc")
}

func main() {
	testa()
	testb(20)
	testc()

	// aaa
	// panic: runtime error: index out of range

}
