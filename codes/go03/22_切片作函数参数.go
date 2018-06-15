package main

import "fmt"
import "math/rand"
import "time"

func initData(s []int) {
	// 设置种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		// 100以内的随机数
		s[i] = rand.Intn(100)
	}
}

func bubbleSort(s []int) {
	n := len(s)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	n := 10

	// 创建一个切片，len为n
	s := make([]int, n)

	initData(s)
	fmt.Println("排序前 -> s = ", s)
	// 输出：排序前 -> s =  [6 10 69 2 97 31 70 76 5 87]

	bubbleSort(s)
	fmt.Println("排序后：s = ", s)
	// 输出：排序后 -> s =  [2 5 6 10 31 69 70 76 87 97]
}
