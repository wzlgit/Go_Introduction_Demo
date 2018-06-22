package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 abcdefg 1.23 7. 8.99 1sdf 6.666"

	// 解析正则表达式，+匹配前一个字符的1次或多次
	reg := regexp.MustCompile(`\d+.\d+`)
	if reg == nil {
		fmt.Println("MustCompile error")
		return
	}

	// 提取关键信息
	//	result := reg.FindAllString(buf, -1)
	// result =  [3.1 4 5 1.2 3 7 8.9 9 1 6.6]

	result := reg.FindAllStringSubmatch(buf, -1)
	// result =  [[3.14] [567] [1.23] [8.99] [6.666]]

	fmt.Println("result = ", result)
}
