package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var m int = 100

var lock sync.Mutex

var wt sync.WaitGroup

func add1() {
	defer wt.Done()
	lock.Lock()
	m = m + 1
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub1() {
	defer wt.Done()
	lock.Lock()
	m = m - 1
	time.Sleep(time.Millisecond * 2)
	lock.Unlock()
}

func GoMutex() {

	// 除了使用channel同步外 还可以使用Mutex互斥锁的方式实现同步
	// TODO：这个demo 感觉有问题 锁去除后 m 同样是100 没测出来读取m有问题的情况 以后补充锁的用法

	for i := 0; i < 100; i++ {

		wt.Add(1)
		go add1()

		wt.Add(1)
		go sub1()
	}

	wt.Wait()
	fmt.Printf("m: %v\n", m)
}
