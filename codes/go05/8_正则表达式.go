package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"

	// 解析规则：解析正则表达式成功则返回解析器
	//	reg1 := regexp.MustCompile(`a.c`)
	//	reg1 := regexp.MustCompile(`a[0-9]c`)
	reg1 := regexp.MustCompile(`a\dc`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	result := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)

	//	result =  [[abc] [azc] [a7c] [aac] [a9c]]
	//  result =  [[a7c] [a9c]]
}
