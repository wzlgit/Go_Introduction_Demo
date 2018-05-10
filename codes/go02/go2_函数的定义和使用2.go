package main

import "fmt"

func myFunc1() int {
  return 88;
}

// Go推荐写法
func myFunc2() (result int) {
  result = 88
  return
}

func main() {
  var a int
  a = myFunc1()
  fmt.Println("a = ", a)
  // output-> a =  88

  b := myFunc1()
  fmt.Println("b = ", b)
  // output-> b =  88

  c := myFunc2()
  fmt.Println("c = ", c)
  // output-> c =  88
}
