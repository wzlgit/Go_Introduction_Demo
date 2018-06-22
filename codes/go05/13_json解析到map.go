package main

import (
	"encoding/json"
	"fmt"
)

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

	// 创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("m = %+v\n", m)
	// m = map[company:google  subjects:[Go Python C++ Test] isok:true price:666.666]

	var str string
	// 类型断言
	for key, value := range m {
		fmt.Printf("%v ---> %v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string, value = %s\n", key, str)
		case bool:
			fmt.Printf("map[%s]的值类型为bool, value = %v\n", key, str)
		case float64:
			fmt.Printf("map[%s]的值类型为float64, value = %v\n", key, str)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为interface{}, value = %v\n", key, str)
		}
	}

	// m = map[company:google  subjects:[Go Python C++ Test] isok:true price:666.666]
	// company ---> google
	// map[company]的值类型为string, value = google
	// subjects ---> [Go Python C++ Test]
	// map[subjects]的值类型为interface{}, value = google
	// isok ---> true
	// map[isok]的值类型为bool, value = google
	// price ---> 666.666
	// map[price]的值类型为float64, value = google
}
