package main

import "fmt"

func main() {
	var p *int
	p = nil
	fmt.Println("p = ", p)

	// runtime error: invalid memory address or nil pointer dereference
	*p = 666
}
