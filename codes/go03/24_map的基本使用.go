package main

import "fmt"

func main() {
	// 定义一个变量，类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	// 输出：m1 =  map[]
	fmt.Println("len = ", len(m1))
	// 对于map，只有len函数，没有cap函数
	// 输出：len =  0

	// 可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	// 可以通过make创建，指定长度，只指定了容量，但是里面却是一个数据也没有
	m3 := make(map[int]string, 2)
	m3[1] = "wenzil"
	m3[2] = "test"
	m3[3] = "go"

	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))
	// 输出：m3 =  map[1:wenzil 2:test 3:go]
	// 		len =  3

	// 键值是唯一的
	m4 := map[int]string{1: "wenzil", 2: "test", 3: "go"}
	fmt.Println("m4 = ", m4)
	// 输出结果如上
}
