package main

import "fmt"

func main()  {
  a, b := 10, 20
  defer func(x int) {
    fmt.Println("defer:", x, b)
  }(a)

  a += 10
  b += 100
  fmt.Printf("a = %d, b =%d\n", a, b)
  /**
  output->
        a = 20, b =120
        defer: 10 120
  */
}
