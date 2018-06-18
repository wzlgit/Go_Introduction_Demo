package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

// 接收者为普通变量，非指针，值语义，一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoValue &p = %p\n", &p)
}

// 接收者为指针变量，引用传递
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoValue p = %p\n", p)
}

func main() {
	s1 := Person{"wenzil", 'm', 20}
	fmt.Printf("&s1 = %p\n", &s1)

	// 值语义
	//	s1.SetInfoValue("wenzil", 'm', 18)
	//	fmt.Println("s1 = ", s1)

	// 引用语义
	(&s1).SetInfoPointer("wenzil", 'm', 88)
	fmt.Println("s1 = ", s1)
}
