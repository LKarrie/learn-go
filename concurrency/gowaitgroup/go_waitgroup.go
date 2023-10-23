package gowaitgroup

import "sync"

var wg sync.WaitGroup

func hello(i int) {
	// goroutine结束就登记 -1
	defer wg.Done()
	println("Hello Goroutine!", i)
}

func GoWaitGroup() {
	for i := 0; i < 10; i++ {
		// goroutine结束就登记 +1
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()

	// 如果去除wg 程序结束会有一些没跑完的goroutine无法输出
}
