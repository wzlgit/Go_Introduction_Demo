package main

import "fmt"

func test1() int {
  var a int   // 函数被调用时，a才分配空间，初始化为0
  a++
  return a * a  // 函数调用完毕，a会自动释放
}

func main1()  {
  fmt.Println(test1())
  fmt.Println(test1())
  fmt.Println(test1())
  /**
  * output->
            1
            1
            1
  *
  */
}

// 函数的返回值是一个匿名函数，返回一个函数类型
func test2() func() int {
  var a int
  return func () int  {
    a++
    return a * a
  }
}

func main()  {
  // 返回值为一个匿名函数，返回一个函数类型，通过t来接收返回的匿名函数，t()来调用匿名函数
  // 闭包不关心这些捕获了的变量和常量是否超出了作用域，只要闭包还在使用它，这些变量一直都存在
  t := test2()
  fmt.Println(t())
  fmt.Println(t())
  fmt.Println(t())
  /**
  * output->
            1
            4
            9
  *
  */
}
