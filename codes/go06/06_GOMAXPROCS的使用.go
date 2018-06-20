package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 指定并行计算的CPU核数的最大值
	n := runtime.GOMAXPROCS(2)
	fmt.Println("n = ", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
