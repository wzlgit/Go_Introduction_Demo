package main

import (
	"fmt"
	"os"
)

func main() {
	//	os.Stdout.Close() // 关闭后，无法输出
	fmt.Println("are u ok?") // 往标准输出设备（屏幕）写内容

	// 标准设备文件，默认已经打开，用户可以直接使用
	os.Stdout.WriteString("are u ok?")

	var a int
	fmt.Println("请输入a:")
	fmt.Scan(&a) // 从标准输入设备中读取内容，放到变量a中
	fmt.Println("a = ", a)

	//	are u ok?
	//  are u ok?
}
