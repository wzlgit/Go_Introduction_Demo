package main

import "fmt"

/**

* output-> 输入a的值：10
*           a =  10
*/
func main()  {
  var a int
  fmt.Printf("输入a的值：")

  // fmt.Scanf("%d", &a)
  fmt.Scan(&a)
  fmt.Println("a = ", a)
}
