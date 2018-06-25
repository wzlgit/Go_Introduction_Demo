package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	// 创建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 使用完毕，需要关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		// "i = %d\n"这个字符串存储在buf中
		buf = fmt.Sprintf("i = %d\n", i)
		_, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
		}
	}
}

func ReadFile(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 使用完毕，需要关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) // 2k大小
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
}

func ReadFileLine(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 使用完毕，需要关闭文件
	defer f.Close()

	// 新建一个缓冲区，把内容放入缓冲区
	r := bufio.NewReader(f)
	for {
		// 遇到"\n"结束读取，但是"\n"也读取进来了
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { // 文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}
}

func main() {
	path := "./demo.txt"
	//	WriteFile(path)
	//	ReadFile(path)
	ReadFileLine(path)
}
