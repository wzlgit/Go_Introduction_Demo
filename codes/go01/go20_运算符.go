package main

import "fmt"

func main() {
  fmt.Println("4 > 3：", 4 > 3)
  fmt.Println("4 != 3：", 4 != 3)

  fmt.Println("!(4 > 3)：", !(4 > 3))
  fmt.Println("!(4 != 3)：", !(4 != 3))

  /**
  * output->
      4 > 3： true
      4 != 3： true
      !(4 > 3)： false
      !(4 != 3)： false
  */
  fmt.Println("(4 > 3) && (5 > 4)", (4 > 3) && (5 > 4))
  fmt.Println("(4 > 3) || (5 < 4)", (4 > 3) || (5 < 4))
  /**
  * output->
          4 > 3 && 5 > 4 true
          4 > 3 || 5 < 4 true
  */
}
