package test

import "fmt"

// 如果首字母小写，只能在同一个包里使用
type Student struct {
	Id int // 首字母大写，才能被外部访问
}

func myFunc() {
	fmt.Println("this is myFunc")
}

func MyFunc() {
	fmt.Println("this is MyFunc")
}
