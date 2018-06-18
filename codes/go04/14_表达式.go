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

	// 方法值  f := p.SetInfoPointer()  // 隐藏了接收者
	f := (*Person).SetInfoPointer
	f(&p) // 显示把接收者传递过去 -> p.SetInfoPointer()

	f2 := (Person).SetInfoValue
	f2(p) // 显示把接收者传递过去 -> p.SetInfoValue

	//	main: 0xc42000a080, {wenzil 109 18}
	// SetInfoPointer: 0xc42000c030, &{wenzil 109 18}
	// SetInfoValue: 0xc42000a0e0, {wenzil 109 18}
}
