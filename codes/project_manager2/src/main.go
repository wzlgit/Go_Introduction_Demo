package main

import (
	"calc"
	"fmt"
)

// 会自动先执行init函数，再执行main函数
func init() {
	fmt.Println("this is init func")
}

func main() {
	a := calc.Add(10, 20)
	fmt.Println("a = ", a)

	fmt.Println("result = ", calc.Minus(10, 5))
}

// export GOPATH=$GOPATH:.../project_manager2/src
// 终端进入到src目录执行如下命令
// go build main.go
// go run main.go
