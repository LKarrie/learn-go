package concurrency

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func GoSelect() {
	// select 是 Go 中的一个控制结构，类似switch语句， 用于处理异步IO操作
	// select 会监听 case 语句中 channel 的读写操作 当case中channel读写操作为非阻塞（即能读能写）将触发相应动作
	// select 中的case语句必须是一个 channel 操作
	// select 中的default子句总是可运行的

	// 如果有多个case都可以运行 select会随机公平地选出一个执行 其他不会执行
	// 如果没有可运行的case 且有default语句 那么就会执行default动作
	// 如果没有可运行的case 没有default语句 select将阻塞 直到某个case通信可以运行

	fmt.Println("测试代码注释了")
	time.Sleep(time.Second)
	// go func() {
	// 	chanInt <- 100
	// 	chanStr <- "hello"
	// 	// 不关闭 会一直阻塞
	// 	close(chanInt)
	// 	close(chanStr)
	// }()

	// for {
	// 	select {
	// 	case r := <-chanInt:
	// 		fmt.Printf("chanInt: %v\n", r)
	// 	case r := <-chanStr:
	// 		fmt.Printf("chanStr: %v\n", r)
	// 	default:
	// 		fmt.Println("default...")
	// 	}
	// 	time.Sleep(time.Second)
	// }

}
