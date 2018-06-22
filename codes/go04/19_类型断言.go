package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1          // int
	i[1] = "hello go" // string
	i[2] = Student{"wenzil", 666}

	// 类型查询，类型断言
	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int，内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string，内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student，内容为name = %s，id = %d\n", index, value.name, value.id)
		}
	}

	//	x[0] 类型为int，内容为1
	//  x[1] 类型为string，内容为hello go
	//  x[2] 类型为Student，内容为name = wenzil，id = 666
}
