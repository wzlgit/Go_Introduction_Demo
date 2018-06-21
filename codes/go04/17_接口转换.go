package main

import "fmt"

type Humaner interface { //子集
	sayHi()
}

type Personer interface { // 超集
	Humaner
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

func (temp *Student) sayHi() {
	fmt.Printf("Student[%s, %d] sayHi\n", temp.name, temp.id)
}

func (temp *Student) sing(lrc string) {
	fmt.Println("Student在唱着：", lrc)
}

func main() {
	// 超集可以转换为子集，反过来不可以
	var iPro Personer // 超集
	var i Humaner     // 子集
	iPro = &Student{"wenzil", 666}

	//	iPro = i	// err
	i = iPro
	i.sayHi()
	// Student[wenzil, 666] sayHi
}
