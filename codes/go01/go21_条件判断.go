package main

import "fmt"

func main() {
  s := "go"
  b := "go"
  if s == b {
    fmt.Println("s == b")
  }
  // output-> s == b

  // if支持一个初始化语句，初始化语句和条件判断，以分号分隔
  if a := 8; a == 10 {
      fmt.Println("a == 10")
  } else if a > 10 {
    fmt.Println("a > 10")
  } else {
    fmt.Println("a < 10")
  }
  // output-> a == 10

  sex := "man"
  switch sex {
  // switch sex := "man"; sex { // 以上两行可以合并为一行，支持一个初始化语句，以分号分隔开
  case "man":
    fmt.Println("性别男")
    // break  ->  默认包含break，写了break也是一样会跳出case条件判断
    // fallthrough -> 不跳出case条件判断，输出“性别男\n性别女”
  case "woman":
    fmt.Println("性别女")
  default:
    fmt.Println("性别不详")
  }
  // output-> 性别男

  // 另外一种用法
  score := 88
  switch  { // switch不添加条件判断
  case score > 90:  // case添加条件判断
    fmt.Println("优秀")
  case score > 80:
    fmt.Println("良好")
  case score > 60:
    fmt.Println("及格")
  default:
    fmt.Println("不及格")
  }
  // output-> 良好
}
