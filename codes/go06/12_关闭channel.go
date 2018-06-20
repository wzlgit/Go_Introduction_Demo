package main

import "fmt"

func main() {
	ch := make(chan int) // 新建一个无缓存的channel

	// 新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // 往通道写数据
		}
		// 不需要再写数据时，关闭channel
		close(ch)
		//		ch <- 666  // 关闭channel后无法再发送数据
	}()

	for num := range ch {
		fmt.Println("num = ", num)
	}

	// num =  0
	// num =  1
	// num =  2
	// num =  3
	// num =  4
}

func main01() {
	ch := make(chan int) // 新建一个无缓存的channel

	// 新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // 往通道写数据
		}
		// 不需要再写数据时，关闭channel
		close(ch)
		//		ch <- 666  // 关闭channel后无法再发送数据
	}()

	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else { // 管道关闭
			break
		}
	}

	// num =  0
	// num =  1
	// num =  2
	// num =  3
	// num =  4
}
