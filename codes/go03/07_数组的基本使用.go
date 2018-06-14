package main

import "fmt"

func main() {
	// 定义一个数组
	var a [10]int
	var b [5]int

	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))
	// 输出：len(a) = 10, len(b) = 5

	// 注意：定义数组时，指定的数组元素个数必须是常量
	//	n := 10
	// 报错：non-constant array bound n
	//	var c [n]int

	// 下标可以是变量或者常量
	i := 1
	a[i] = 2

	// 给每个元素赋值
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	// 打印：第一个返回数组下标，第二个返回元素
	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}

	/*
		* 输出：
			a[0] = 1
			a[1] = 2
			a[2] = 3
			a[3] = 4
			a[4] = 5
			a[5] = 6
			a[6] = 7
			a[7] = 8
			a[8] = 9
			a[9] = 10
	*/
}
