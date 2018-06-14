package main

import "fmt"

func modify(a [5]int) {
	a[0] = 666
	fmt.Println("main: a =", a)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	modify(a)
	fmt.Println("main: a =", a)

	// 输出：
	// main: a = [666 2 3 4 5]
	// main: a = [1 2 3 4 5]
}
