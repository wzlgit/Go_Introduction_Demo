package main

import "fmt"

func main() {
	// 作用域：变量作用的范围

	if flag := 3; flag == 3 {
		fmt.Fprintln("flag = ", flag)
	}

	flag = 4
	// output-> undefined: flag
}
