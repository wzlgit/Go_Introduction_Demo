package main

import (
	"fmt"
	"net/http"
	//	"os"
	"regexp"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	// 开始爬取网页内容
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}

		result += string(buf[:n])
	}
	return
}

func SpiderPage(i int) {
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页: %s\n", i, url)

	// 开始爬取网页内容
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}
	//	fmt.Println("r = ", result)
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	joyUrls := re.FindAllStringSubmatch(result, -1)
	//	fmt.Println("joyUrls = ", joyUrls)

	for _, data := range joyUrls {
		fmt.Println("url = ", data[1])
	}
}

func DoWork(start, end int) {
	fmt.Printf("准备爬取第%页到%d页的网址\n", start, end)

	for i := start; i <= end; i++ {
		SpiderPage(i)
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页( >= 1) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	DoWork(start, end)
}
