package main

import "fmt"
import "math/rand"
import "time"

func createNum(p *int) {
	// 设置种子
	rand.Seed(time.Now().UnixNano())

	var num int

	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}

	//	fmt.Println("num = ", num)
	*p = num
}

func getNum(s []int, num int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func onGame(randSlice []int) {
	var num int
	keySlice := make([]int, 4)

	for {
		for {
			fmt.Printf("请输入一个4位数：")
			fmt.Scan(&num)
			// 999 < num < 10000
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("输入的数不符合要求")
		}
		//	fmt.Println("num = ", num)

		getNum(keySlice, num)
		//	fmt.Println("keySlice = ", keySlice)

		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了一点\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了一点\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}

		if n == 4 {
			fmt.Println("全部猜对了！！！")
			break
		}
	}
}

func main() {
	var randNum int

	// 产生一个4位的随机数
	createNum(&randNum)
	fmt.Println("randNum = ", randNum)

	// 保存这个4位数的每一位
	randSlice := make([]int, 4)
	getNum(randSlice, randNum)
	fmt.Println("randSlice = ", randSlice)

	onGame(randSlice)
}
