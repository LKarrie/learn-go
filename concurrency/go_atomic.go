package concurrency

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func GoAtomic() {
	// 除了使用Mutex包中的锁实现 协程同步

	// 使用原子变量 也可以
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("i: %v\n", i)

	// atomic 提供的原子操作能够确保任意时刻只有一个goroutine对变量进行操作 避免使用大量的锁操作
	// 增减操作
	// func AddInt32(addr *int32, delta in32) (new int32)
	// ...
	// 载入操作 保证读取时 任何其他CPU操作都无法对变量进行读写
	// func LoadInt32(addr *int32)(val int32)
	// ...
	// 比较并交换 该操作在进行交换前首先确保变量的值未被改变 即保持 old 所记录的值
	// func CompareAndSwapInt32(addr *int32,old,new int32)(swapped bool)
	// ...
	// 交换 不论变量的旧值是否改变 直接赋予新值并返回替换的值
	// func SwapInt32(addr *int32,new int32)(old int32)
	// ...
	// 存储 保证写变量的原子性 避免其他操作读到了修改变量过程中的脏数据
	// func StoreInt32(addr *int32,val int32)
	// ...

}
