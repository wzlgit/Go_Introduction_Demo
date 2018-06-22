package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := `
	{
	"company": "google ",
	"subjects": [
		"Go",
		"Python",
		"C++",
		"Test"
	],
	"isok": true,
	"price": 666.666
	}	
	`

	var temp IT // 定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &temp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//	fmt.Println("temp = ", temp)
	// temp =  {google  [Go Python C++ Test] true 666.666}

	fmt.Printf("temp = %+v\n", temp)
	// temp = {Company:google  Subjects:[Go Python C++ Test] IsOk:true Price:666.666}

	type IT2 struct {
		Company string `json:"company"`
	}

	var temp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &temp2)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("temp2 = %+v\n", temp2)
	// temp2 = {Company:google }
}
