package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args // 获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage: xxx src File dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	// 只读方式打开源文件
	sF, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 新建目的文件
	dF, err := os.Create(dstFileName)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 操作完毕，需要关闭文件
	defer sF.Close()
	defer dF.Close()

	// 核心处理，从源文件读取内容，往目的文件写
	buf := make([]byte, 4*1024) // 4k大写临时缓冲区
	for {
		n, err := sF.Read(buf)
		if err != nil {
			if err == io.EOF { // 文件读取完毕
				break
			}
			fmt.Println("err = ", err)
		}
		dF.Write((buf[:n]))
	}

	// 运行命令：【
	// go build 17_拷贝文件.go
	// ./17_拷贝文件 demo.txt demo2.txt
}
