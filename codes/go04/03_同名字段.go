package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
	name   string // 和Person同名了
}

func main() {
	var s Student

	// 默认规则（就近原则），如果能在本作用域找到此成员，就操作此成员
	s.name = "wenzil"
	s.sex = 'm'
	s.age = 18
	s.addr = "jingzhou"

	s.Person.name = "test"
	fmt.Printf("s = %+v\n", s)
	// s = {Person:{name:test sex:109 age:18} id:0 addr:jingzhou name:wenzil}
}
