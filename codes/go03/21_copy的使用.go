package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}
	destSlice := []int{6, 6, 6, 6, 6}

	copy(destSlice, srcSlice)
	fmt.Println("destSlice = ", destSlice)
	// 输出：destSlice =  [1 2 6 6 6]
}
