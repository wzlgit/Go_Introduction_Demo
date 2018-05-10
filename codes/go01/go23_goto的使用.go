package main

import "fmt"

func main() {
  fmt.Println("start")
  fmt.Println("one")
  fmt.Println("two")
  goto End
  // 无条件跳转，跳转到定义为"End"的代码段执行，不建议使用，因为它很难跟踪程序的控制流程
  fmt.Println("three")

End:
  fmt.Println("end")
}

/**
* output->
*         start
          one
          two
          end
*
/
