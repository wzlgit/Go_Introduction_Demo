package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓存的channel
	ch := make(chan int, 0)

	// len(ch)为缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程：i = ", i)
			ch <- i // 往通道中写内容[
			fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch // 读取通道中内容，没有内容前，阻塞
		fmt.Println("num = ", num)
	}

	// 每次运行结果都不一样，打印"num = xx"之前会阻塞
	//len(ch) = 0, cap(ch) = 0
	// 子协程：i =  0
	// len(ch) = 0, cap(ch) = 0
	// num =  0
	// 子协程：i =  1
	// len(ch) = 0, cap(ch) = 0
	// 子协程：i =  2
	// num =  1
	// num =  2
}
