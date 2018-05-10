package main

import "fmt"

func main() {
  type bigint int64

  var a bigint
  fmt.Printf("a type is %T\n", a)
  // output-> a type is main.bigint

  type (
    long int64
    char byte
  )

  var b long = 10
  var c char = 'c'
  fmt.Printf("b = %d, c = %c\n", b, c)
  // output-> b = 10, c = c
}
