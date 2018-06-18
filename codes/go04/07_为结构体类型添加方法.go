package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (temp Person) PrintInfo() {
	fmt.Println("temp = ", temp)
}

func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {
	// 定义同时初始化
	p := Person{"wenzil", 'm', 18}
	p.PrintInfo()
	// temp =  {wenzil 109 18}

	var p2 Person
	(&p2).SetInfo("test", 'f', 22)
	p2.PrintInfo()
	// temp =  {test 102 22}]]]]
}
