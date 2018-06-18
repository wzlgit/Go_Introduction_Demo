package main

import "fmt"

type mystr string // 自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	mystr
}

func main() {
	s := Student{Person{"wenzil", 'm', 18}, 666, "test..."}
	fmt.Printf("s = %+v\n", s)
	// s = {Person:{name:wenzil sex:109 age:18} id:666 mystr:test...}

	s.Person = Person{"go", 'm', 22}
	fmt.Printf("s = %+v\n", s)
	// s = {Person:{name:go sex:109 age:22} id:666 mystr:test...}
}
