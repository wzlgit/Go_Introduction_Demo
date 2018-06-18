package main

import "test"
import "fmt"

func main() {
	test.MyFunc()
	// this is MyFunc

	var s test.Student
	s.Id = 666
	fmt.Println("s = ", s)
	// s =  {0}
}
