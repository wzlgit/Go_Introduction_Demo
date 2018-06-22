package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("helloworld", "hello"))
	fmt.Println(strings.Contains("helloworld", "go"))

	s := []string{"welcome", "to", "hello", "world"}
	buf := strings.Join(s, "x")
	fmt.Println("buf = ", buf)

	fmt.Println(strings.Index("hello", "o"))
	// 不包含子串返回-1
	fmt.Println(strings.Index("hello", "a"))

	buf = strings.Repeat("go", 3)
	fmt.Println("buf = ", buf)

	// Split：以指定的分隔符拆分
	buf = "welcome,to,hello,world"
	s2 := strings.Split(buf, ",")
	fmt.Println("s2 = ", s2)

	// Trim：去掉两头的字符
	buf = strings.Trim("   are u ok?   ", " ")
	fmt.Printf("buf = #%s#\n", buf)

	// Fields：去掉空格，把元素放入切片中
	s3 := strings.Fields("   are u ok?   ")
	for i, data := range s3 {
		fmt.Println(i, " : ", data)
	}

	//	true
	//	false
	//	buf = welcomextoxhelloxworld
	// 4
	// -1
	// buf =  gogogo
	// s2 =  [welcome to hello world]
	// buf = #are u ok?#
	//	0  :  are
	//  1  :  u
	//  2  :  ok?
}
