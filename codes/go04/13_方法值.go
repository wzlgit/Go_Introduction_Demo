package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", &p, p)
}

func main() {
	p := Person{"wenzil", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	// 传统调用方法
	p.SetInfoPointer()
	// SetInfoPointer: 0xc420090020, &{wenzil 109 18}

	// 保存方式入口地址
	// 这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者
	pFunc := p.SetInfoPointer
	pFunc()
	// SetInfoPointer: 0xc420090028, &{wenzil 109 18}

	vFunc := p.SetInfoValue()
	vFunc() // 等价于p.SetInfoValue()
}
