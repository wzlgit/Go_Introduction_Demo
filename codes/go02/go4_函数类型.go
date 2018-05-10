package main

import "fmt"

func add(a, b int) int {
  return a + b
}

func minus(a, b int) int {
  return a - b
}

// 函数类型
type FuncType func(int, int) int

func main()  {
  var result int
  result = add(1, 1)
  fmt.Println("result = ", result)
  // output-> result =  2

  var fTest FuncType
  fTest = add
  result =  fTest(10, 20)
  fmt.Println("result = ", result)
  // output-> result =  30

  fTest = minus
  result =  fTest(10, 4)
  fmt.Println("result = ", result)
  // output-> result =  6
}
