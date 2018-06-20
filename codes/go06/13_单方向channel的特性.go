package main

func main() {
	// 创建一个channel，双向的
	ch := make(chan int)

	// 双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch // 只能写，不能读
	var readCh <-chan int = ch  // 只能读，不能写

	writeCh <- 666 // 写
	//	<-writeCh
	// invalid operation: <-writeCh (receive from send-only type chan<- int)
	// 不要运行，编译go即可，命令为："go build 13_单方向channel的特性.go"

	<-readCh // 读
	//	readCh <- 666 // 写
	// invalid operation: readCh <- 666 (send to receive-only type <-chan int)

	// 单向无法转换为双向
	var ch2 chan int = writeCh
	// cannot use writeCh (type chan<- int) as type chan int in assignment
}
