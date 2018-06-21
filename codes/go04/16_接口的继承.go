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
	// 定义一个接口类型的变量
	var i Personer
	s := &Student{"wenzil", 666}
	i = s

	i.sayHi()
	i.sing("好汉歌")

	//	Student[wenzil, 666] sayHi
	//  Student在唱着： 好汉歌
}
