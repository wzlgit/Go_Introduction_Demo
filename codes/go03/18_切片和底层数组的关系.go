package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 新切片，从a[2]开始，取3个元素
	s1 := array[2:5]
	s1[1] = 666
	fmt.Println("s1 = ", s1)
	// 输出：s1 =  [2 666 4]
	fmt.Println("array = ", array)
	// 输出：array =  [0 1 2 666 4 5 6 7 8 9]

	// 另外一个新切片
	s2 := s1[2:7]
	s2[2] = 777
	fmt.Println("s2 = ", s2)
	// 输出：s2 =  [4 5 777 7 8]
	fmt.Println("array = ", array)
	// 输出：array =  [0 1 2 666 4 5 777 7 8 9]

}
