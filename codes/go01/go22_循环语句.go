package main

import "fmt"

func main() {
  sum := 0
  for i := 1; i <= 5; i++ {
    sum += i
  }
  fmt.Println("sum = ", sum)
  // output-> sum =  15

  str := "abc"
  for i := 0;i < len(str); i++ {
    fmt.Printf("str[%d] = %c\n", i, str[i])
  }
  /**
  * output->
  *         str[0] = a
  *         str[1] = b
  *         str[2] = c
  */
  // i返回的是元素的位置，data返回的是元素的值
  for i, data := range str {
    fmt.Printf("str[%d] = %c\n", i, data)
  }
  // output-> 同上

  // 丢弃第二个返回值，可以用下划线"_"表示
  // 或者直接用"for i := range str"表示
  for i, _ := range str {
    fmt.Printf("str[%d] = %c\n", i, str[i])
  }
  // output-> 同上

  for i := 0; i < 5; i++ {
    if i == 3 {
      break  // 跳出整个循环
      // continue  // 跳出本次循环
    }
    fmt.Println("i = ", i)
  }
  /**
  * 使用了break的output->
                      i =  0
                      i =  1
                      i =  2
  * 使用了continue的output->
                        i =  0
                        i =  1
                        i =  2
                        i =  4
  */
}
