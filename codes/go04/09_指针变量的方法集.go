package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	// 结构体变量是一个指针变量，它能够调用那些方法，这些方法就是一个集合，简称方法集
	p := &Person{"wenzil", 'm', 18}
	p.SetInfoPointer()
	p.SetInfoValue()

	// SetInfoPointer
	// SetInfoValue
}
