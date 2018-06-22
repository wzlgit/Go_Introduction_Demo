package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量名首字母必须大写
//type IT struct {
//	Company  string
//	Subjects []string
//	IsOk     bool
//	Price    float64
//}

type IT struct {
	Company  string   `json:"-"`        // 此字段不会输出到屏幕
	Subjects []string `json:"subjects"` // 二次编码
	IsOk     bool     `json:",string"`
	Price    float64  `json:",string"`
}

func main() {
	// 定义一个结构体变量，同时初始化
	s := IT{"google", []string{"Go", "Java", "Python", "Test"}, true, 66.666}

	// 编码：根据内容生成json文本
	//	buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ") // 格式化编码，代码缩进
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
	// buf =  {"Company":"google","Subjects":["Go","Java","Python","Test"],"IsOk":true,"Price":66.666}
}
