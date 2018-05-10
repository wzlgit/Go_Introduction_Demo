package main

import "fmt"

func main() {
  var flag bool
  flag = true
  fmt.Printf("flag = %t\n", flag)
  // output-> flag = true

  // bool类型不能转换为int
  // fmt.Printf("flag = %d\n", int(flag))
  // ./go18_类型转换.go:11:32: cannot convert flag (type bool) to type int

  // 整型不能转换为bool
  // flag = bool(1)

  var ch byte
  ch = 'a'
  var chInt int
  chInt = int(ch)
  fmt.Println("chInt = ", chInt)
  // output-> chInt =  97
}
