package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个有缓存的channel
	ch := make(chan int, 3)

	// len(ch)为缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 往通道中写内容[
			fmt.Printf("子协程[%d] : len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch // 读取通道中内容，没有内容前，阻塞
		fmt.Println("num = ", num)
	}

	/*
		* 每次运行结果都不一样，可以看作是边读边写
		* len(ch) = 3时缓冲区满了，然后阻塞，读取通道内容才能再写入数据
		*
			len(ch) = 0, cap(ch) = 3
			子协程[0] : len(ch) = 0, cap(ch) = 3
			子协程[1] : len(ch) = 1, cap(ch) = 3
			子协程[2] : len(ch) = 2, cap(ch) = 3
			子协程[3] : len(ch) = 3, cap(ch) = 3
		   num =  0
		   num =  1
		   num =  2
		   num =  3
		   num =  4
		   子协程[4] : len(ch) = 0, cap(ch) = 3
		   子协程[5] : len(ch) = 0, cap(ch) = 3
		   子协程[6] : len(ch) = 1, cap(ch) = 3
		   子协程[7] : len(ch) = 2, cap(ch) = 3
		   子协程[8] : len(ch) = 3, cap(ch) = 3
		   num =  5
		   num =  6
		   num =  7
		   num =  8
		   num =  9
	*/

}
