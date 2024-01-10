package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func GoChannel() {
	// Go提供通道的机制，用于在gorotine之间共享数据
	// 需要在声明通道时指定数据类型
	// 数据在通道上传递时，任何时间内只有一个goroutine可以访问数据项
	// 根据数据交换的行为有两种通道类型，无缓冲通道和有缓冲通道
	// 无缓冲通道用于执行 goroutine 之间的同步通信
	// 缓冲通道用于执行异步通信
	// 通道由 make 函数创建 该函数指定 chan 关键字和通道的元素类型

	// 整型无缓冲通道
	// unbuffered := make(chan int)
	// 整型有缓冲通道
	// buffered := make(chan int,10)

	// 将值发送到通道的代码块需要使用 <- 运算符
	// goroutine1 := make(chan string,5)
	// 通过通道发送字符串
	// gotoutine1 <- "123"

	// 通过通道接收字符串
	// data := <-gotoutine1

	// 通道发送和接收的特性
	// 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的
	// 发送操作和接收操作中对元素的处理都是不可分割的
	// 发送操作在完全完成之前会被阻塞，接收操作也一样

	// 从通道接收
	defer close(values)
	go send()
	println("wait...")
	value := <-values
	println("receive: %v\n", value)
	println("end")

	// channel遍历 for 循环+if
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Printf("data: %v\n", data)
		} else {
			fmt.Printf("ok: %v\n", ok)
			break
		}
	}
	// channel遍历 for range
	cc := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			cc <- i
		}
		close(cc)
	}()
	for v := range cc {
		fmt.Printf("v: %v\n", v)
	}

	// 如果通道关闭 多读写少 没有了就是默认值 例如 int 就是0 如果没有关闭就会死锁
	time.Sleep(time.Millisecond * 200)
}

func send() {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(10)
	println("send: %v\n", value)

	values <- value
}
