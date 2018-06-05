package main

import "fmt"

func main()  {
  // defer延迟调用，main函数结束前调用
  defer fmt.Println("defer...")
  fmt.Println("main")
  /**
  * output->
          main
          defer...
  *
  */

  // 如果一个函数有多个defer语句，它们会已LIFO（后进先出）的顺序执行，
  // 哪怕函数或某个延迟调用发生错误，这些调用都会被执行
  defer fmt.Println("1")
  defer fmt.Println("2")
  defer test(0)
  defer fmt.Println("3")
  /**
  * output->
          3
          2
          1
          defer...
          panic: runtime error: integer divide by zero


  *
  */
}

func test(x int)  {
  // x为0时，产生异常
  fmt.Println(10 / x)
}
