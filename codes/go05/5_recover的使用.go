package main

import "fmt"

func testa() {
	fmt.Println("aaa")
}

func testb(x int) {
	// 设置recover
	defer func() {
		//		recover()	// 可以打印panic的错误信息
		//		fmt.Println(recover())
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

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
	// runtime error: index out of range
	// ccc

}
