package main

import "fmt"

// 定义接口类型
type Humaner interface {
	// 方法：只有声明，没有实现，由别的类型（自定义类型）实现
	sayHi()
}

type Student struct {
	name string
	id   int
}

func (temp *Student) sayHi() {
	fmt.Printf("Student[%d, %s] sayHi\n", temp.id, temp.name)
}

type Teacher struct {
	addr  string
	group string
}

func (temp *Teacher) sayHi() {
	fmt.Printf("Teacher[%s, %s] sayHi\n", temp.addr, temp.group)
}

type MyStr string

func (temp *MyStr) sayHi() {
	fmt.Printf("MyStr[%s] sayHi\n", *temp)
}

// 定义一个普通函数，函数的参数为接口类型
// 只有一个函数，可以有不同表现，多态
func WhoSayHi(i Humaner) {
	i.sayHi()
}

func main() {
	s := &Student{"wenzil", 666}
	t := &Teacher{"bj", "go"}
	var str MyStr = "hello world"

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	// 输出结果跟main01一样

	// 创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, i := range x {
		i.sayHi()
	}
}

func main01() {
	// 定义接口类型的变量
	var i Humaner

	// 只要实现了此接口方法的类型，那么这个类型的变量就可以给i赋值
	s := &Student{"wenzil", 666}
	i = s
	i.sayHi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayHi()

	var str MyStr = "hello world"
	i = &str
	i.sayHi()

	// Student[666, wenzil] sayHi
	// Teacher[bj, go] sayHi
	// MyStr[hello world] sayHi
}
