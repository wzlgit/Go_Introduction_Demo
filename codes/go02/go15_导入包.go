package main // 必须

/*
// 方式1
import "fmt"  // 导入包，必须使用，否则编译失败
import "os"

// 方式2
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("this is a test")
	fmt.Println("os.Args = ", os.Args)
}

// 方式3-> 点操作
import . "fmt"	// 调用函数，无需通过包名
import . "os"

// 方式4
// 给包名起别名
import io "fmt"

func main() {
	io.Println("this is a test")
}

// 方式5：忽略此包
// 导入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数
import _ "fmt"

func main() {

}
*/
