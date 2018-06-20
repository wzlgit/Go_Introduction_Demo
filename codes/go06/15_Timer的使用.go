package main

import (
	"fmt"
	"time"
)

// 验证time.NewTimer()，时间到了，只会响应因此
func main() {
	timer := time.NewTimer(2 * time.Second)

	for {
		<-timer.C
		fmt.Println("时间到")
	}

	//	时间到
	// fatal error: all goroutines are asleep - deadlock!
}

func main01() {
	// 创建一个定时器，设置时间为2s，2s后，往time通道写内容（当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	// 2s后，往timer.C写数据，有数据后就可以读取
	t := <-timer.C // channel没有数据前会阻塞
	fmt.Println("t = ", t)

	//	当前时间： 2018-06-20 15:33:45.780647833 +0800 CST m=+0.000418484
	// t =  2018-06-20 15:33:47.780961461 +0800 CST m=+2.000843658
}
