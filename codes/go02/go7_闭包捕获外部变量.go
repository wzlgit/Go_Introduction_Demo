package main

import "fmt"

func main()  {
  a := 2
  str := "hello"

  func ()  {
    a = 123
    str = "hello world"
    fmt.Printf("func->a = %d, str = %s\n", a, str)
  }()

  fmt.Printf("main->a = %d, str = %s\n", a, str)
  /**
  * output->
  *         func->a = 123, str = hello world
  *         main->a = 123, str = hello world
  */
}
