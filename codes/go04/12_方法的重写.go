package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (temp *Person) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", temp.name, temp.sex, temp.age)
}

type Student struct {
	Person
	id   int
	addr string
}

// Student也实现一个方法，这个方法和Person方法同名，这种方法叫重写
func (temp *Student) PrintInfo() {
	fmt.Println("Student: temp = ", temp)
}

func main() {
	s := Student{Person{"wenzil", 'm', 18}, 666, "jingzhou"}
	s.PrintInfo()
	// Student: temp =  &{{wenzil 109 18} 666 jingzhou}

	// 显示调用继承的方法
	s.Person.PrintInfo()
	// name=wenzil, sex=m, age=18
}
