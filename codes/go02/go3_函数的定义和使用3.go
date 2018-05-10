package main

import "fmt"

// 多个返回值
func myFunc1() (int, int, int) {
  return 1, 2, 3
}

func maxAndMin(a, b int) (max, min int) {
  if a > b {
    max = a
    min = b
  } else {
    max = b
    min = a
  }
  return
}

// 递归函数
func fib(n int) uint {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fib(n-1) + fib(n-2)
    }
}

func main()  {
  a, b, c := myFunc1()
  fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
  // output-> a = 1, b = 2, c = 3

  max, min := maxAndMin(10, 20)
  fmt.Printf("max = %d, min = %d\n", max, min)
  // output-> max = 20, min = 10

  n := 10
  fmt.Println(fib(n))
  // output-> 55
}
