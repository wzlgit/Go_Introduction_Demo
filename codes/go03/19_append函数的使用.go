package main

import "fmt"

func main() {
	s1 := []int{}
	showInfo(s1)
	// len = 0, cap = 0
	// s1 =  []

	// 在原切片的末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	showInfo(s1)
	// len = 3, cap = 4
	// s1 =  [1 2 3]

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)
	// s2 =  [1 2 3]
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Println("s2 = ", s2)
	// s2 =  [1 2 3 5 5 5]
}

func showInfo(s []int) {
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	fmt.Println("s1 = ", s)
}
