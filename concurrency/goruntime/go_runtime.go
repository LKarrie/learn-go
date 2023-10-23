package goruntime

import (
	"fmt"
	"runtime"
	"time"
)

func GoRuntime() {
	// 让出CPU时间片，重新等待安排任务
	// runtime.Gosched()
	go show("java")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下 再次分配任务
		// java 会先输出
		runtime.Gosched()
		println("golang")
	}

	// 退出当前协程
	// runtime.Goexit()
	go show2()
	time.Sleep(time.Second)

	// GOMAXPROCS
	println("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

func show(s string) {
	for i := 0; i < 2; i++ {
		println(s)
	}
}

func show2() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			runtime.Goexit()
		}
		fmt.Printf("i: %v\n", i)
	}
}

func a() {
	for i := 1; i < 10; i++ {
		println("A:", i)
	}
}
func b() {
	for i := 1; i < 10; i++ {
		println("B:", i)
	}
}
