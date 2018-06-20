package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时2s，阻塞2s，2s后产生一个时间，往channel写内容
	time.After(2 * time.Second)
	fmt.Println("时间到")
}

func main02() {
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
}

func main01() {
	// 延时2s后打印一句话
	timer := time.NewTimer(2 * time.Second)
	t := <-timer.C
	fmt.Println("t = ", t)
	fmt.Println("时间到")

	// t =  2018-06-20 15:41:01.821323271 +0800 CST m=+2.000698335
	// 时间到
}
