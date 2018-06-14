package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)
	// 输出：a =  [1 2 3 4 5]

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	// 输出：b =  [1 2 3 4 5]

	// 部分初始化，没有初始化的元素，自动赋值为0
	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)
	// 输出：c =  [1 2 3 0 0]

	// 指定某个元素初始化
	d := [5]int{2: 10, 4: 10}
	fmt.Println("d = ", d)
	// 输出：d =  [0 0 10 0 10]
}
