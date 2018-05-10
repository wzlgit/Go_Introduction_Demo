package main

import "fmt"

func main() {
  myFunc1()
  myFunc2(1, 2)
  myFunc3(1, 2, 3)
  test(1, 2, 3, 4)
}

// 函数的定义可以放在main函数后面
func myFunc1() {
  fmt.Println("myFunc1")
}

// 普通参数列表
func myFunc2(a int, b int) {
  var sum int  // 要先定义，否则报错： undefined: sum
  sum = a + b
  fmt.Printf("%d + %d = %d\n", a, b, sum)
}
/**
* myFunc2(1, 2)
* output->
          myFunc1
          1 + 2 = 3
*/

// args ...为不定参数，可以传0或n个参数
func myFunc3(a int, args ...int) {
  fmt.Println("a = ", a)
  fmt.Println("args参数个数为:", len(args))
  for i, data := range args {
    fmt.Printf("args[%d] = %d\n", i, data)
  }
}
/** myFunc3(1, 2, 3)
* output->
            a =  1
            args参数个数为: 2
            args[0] = 1
            args[1] = 2
*
*/

// 报错：不定参数只能放在参数的最后一个参数
// can only use ... with final parameter in list
// func myFunc3(args ...int, a int) {
//
// }

// 不定参数的传递
func myFun4(args ...int) {
  for _ , data := range args {
    fmt.Println("data = ", data)
  }
}

func test(args ...int)  {
  myFun4(args...)
  /**
  * main函数中调用：test(1, 2, 3, 4)
  * output->
            data =  1
            data =  2
            data =  3
            data =  4
  */

  // myFun4(args[:2]...)
  /**
  * 传递的是下标从0到2（不包括2）的参数
  * main函数中调用：test(1, 2, 3, 4)
  * output->
            data =  1
            data =  2
  */

  // myFun4(args[2:]...)
  /**
  * 传递的是下标从2开始到结尾的参数
  * main函数中调用：test(1, 2, 3, 4)
  * output->
            data =  3
            data =  4
  */

}
