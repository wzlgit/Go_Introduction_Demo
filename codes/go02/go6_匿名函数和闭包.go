package main

import "fmt"

func main()  {
  a := 2
  str := "hello"
  // 匿名函数，没有函数名
  func1 := func ()  {
    fmt.Println("a = ", a)
    fmt.Println("str = ", str)
  }
  func1() // 调用匿名函数
  // output-> a =  2
  //         str =  hello

  // 给函数类型器别名，不推荐使用，比较繁琐
  type FuncType func()
  var func2 FuncType
  func2 = func1
  func2()
  // output-> 同上

  // func ()  {
  //   fmt.Println("a = ", a)
  //   fmt.Println("str = ", str)
  // }()
  // output-> 同上

  // 带参数的匿名函数
  func3 := func (i, j int)  {
    fmt.Printf("i = %d, j = %d\n", i, j)
  }
  func3(10, 20)
  // output-> i = 10, j = 20

  // 定义匿名函数，同时调用
  func (i, j int)  {
    fmt.Printf("i = %d, j = %d\n", i, j)
  }(10, 20)
  // output-> i = 10, j = 20

  // 匿名函数，带返回值
  x,y := func (i, j int) (max, min int) {
    if i > j {
      max = i
      min = j
    } else {
      max = j
      min = i
    }
    return
  }(10, 20)
  fmt.Printf("x = %d, y = %d\n", x, y)
  // output-> x = 20, y = 10
}
