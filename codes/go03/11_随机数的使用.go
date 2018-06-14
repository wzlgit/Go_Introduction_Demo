package main

import "fmt"
import "math/rand"
import "time"

func main() {
	// 设置种子，只需要设置一次
	// 如果种子参数一样，每次运行程序产生的随机数都一样
	//	rand.Seed(666)

	// 以当前系统时间作为种子参数
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		// 产生随机数，随机数很大
		//		fmt.Println("rand = ", rand.Int())
		// 限制在100以内的随机数
		fmt.Println("rand = ", rand.Intn(100))
	}
}
