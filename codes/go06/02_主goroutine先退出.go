package main

import (
	"fmt"
	"time"
)

// 主协程退出了，其他子协程也要跟着退出
func main() {
	go func() {

		i := 0
		for {
			i++
			fmt.Println("子 i = ", i)
			time.Sleep(time.Second)

			if i == 2 {
				break
			}
		}

	}()

	i := 0
	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(time.Second)

		if i == 2 {
			break
		}
	}
}
