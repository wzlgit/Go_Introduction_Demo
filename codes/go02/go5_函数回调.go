package main

import "fmt"

type FuncType func(int, int) int

func add(a, b int) int {
  return a + b
}

func minus(a, b int) int {
  return a + b
}


func calc(a, b int, fTest FuncType) (result int) {
  result = fTest(a, b)
  return
}

func main()  {
  result := calc(1, 1, add)
  fmt.Println("result = ", result)
}
// output-> result =  2
