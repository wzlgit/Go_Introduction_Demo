package main

import "fmt"

func main() {
  a:=20
  b:="test"
  c:='t'
  d:=1.23
  /**
  * %T操作变量所属类型
  * output-> int, string, int32, float64
  */
  fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)
  /*
  * %d 整型格式
  * %s 字符串格式
  * %c 字符型格式
  * %f 浮点型格式
  * output-> 20, test, t, 1.230000
  */
  fmt.Printf("%d, %s, %c, %f\n", a, b, c, d)

  /**
   * %v自动匹配格式输出
   * output-> 20, test, 116, 1.23
   */
  fmt.Printf("%v, %v, %v, %v\n", a, b, c, d)
}
