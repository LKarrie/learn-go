package main

import (
	"fmt"
	"learn-go/basic"
	"learn-go/concurrency"
	"learn-go/goos"
)

func main() {
	// 基础
	// gobasic()
	// 并发编程
	// goconcurrency()
	// 标准库
	// os 包
	os()
}

func gobasic() {
	fmt.Printf("------------GoVar------------\n")
	//go中的变量声明和初始化
	//go中的常量
	//go中得iota
	basic.GoVar()

	fmt.Printf("------------GoVarType------------\n")
	//go中的整数浮点数复数
	basic.GoVarType()

	fmt.Printf("------------GoString------------\n")
	//go中的字符字符串和常用方法
	basic.GoString()

	fmt.Printf("------------GoPlaceholder------------\n")
	//go中的占位符
	basic.GoPlaceholder()

	fmt.Printf("------------GoOperator------------\n")
	//go中的运算符
	basic.GoOperator()

	fmt.Printf("------------GoIf------------\n")
	//go中的if if中的goto
	basic.GoIf()

	fmt.Printf("------------GoSwitch------------\n")
	//go中的switch switch中的break和fallthrough
	basic.GoSwitch()

	fmt.Printf("------------GoLoop------------\n")
	//go中的for和range循环 循环中的break和goto 递归
	basic.GoLoop()

	fmt.Printf("------------GoArray------------\n")
	//go中的数组 访问数组 和 遍历数组
	basic.GoArray()

	fmt.Printf("------------GoSlice------------\n")
	//go中的切片 切片的长度和容量 切片的初始化 空切片 切片的遍历
	basic.GoSlice()

	fmt.Printf("------------GoMap------------\n")
	//go中的Map 判断Map键是否存在 遍历Map的key和value
	basic.GoMap()

	fmt.Printf("------------GoFunc------------\n")
	//go中的函数 函数类型 函数的返回值 函数的参数 函数类型和函数变量(函数作为参数) 函数做返回值 匿名函数 闭包
	basic.GoFunc()

	fmt.Printf("------------GoDefer------------\n")
	//go中的defer语句
	basic.GoDefer()

	fmt.Printf("------------GoInit------------\n")
	//go中的init语句
	basic.GoInit()

	fmt.Printf("------------GoPointer------------\n")
	//go中的指针 指向数组的指针
	basic.GoPointer()

	fmt.Printf("------------GoStruct------------\n")
	//go中的结构体 结构体访问 初始化 结构体指针 结构体做参 嵌套结构体（继承）
	basic.GoStruct()

	fmt.Printf("------------GoFunc2------------\n")
	//go中的方法 方法的接收者类型
	basic.GoFunc2()

	fmt.Printf("------------GoInterface------------\n")
	//go中的接口 接口实现 接口接收者类型 接口嵌套
	basic.GoInterface()

	fmt.Printf("------------GoOcp------------\n")
	//go OCP 设计原则（面向对象可复用 开-闭原则 面对扩展开放 面对修改关闭）
	basic.GoOcp()

	fmt.Printf("------------GoExtends------------\n")
	//go 继承 模拟类属性和方法 模拟类构造函数
	basic.GoExtends()
}

func goconcurrency() {
	fmt.Printf("------------GoGoroutines------------\n")
	// 协程
	concurrency.GoGoroutines()
	fmt.Printf("------------GoChannel------------\n")
	// channel通道 channel遍历
	concurrency.GoChannel()
	fmt.Printf("------------GoWaitGroup------------\n")
	// WaitGroup同步
	concurrency.GoWaitGroup()
	fmt.Printf("------------GoRuntime------------\n")
	// runtime包中的一些方法
	concurrency.GoRuntime()
	fmt.Printf("------------GoMutex------------\n")
	// mutex包中lock实现同步
	concurrency.GoMutex()
	fmt.Printf("------------GoSelect------------\n")
	// GoSelect 类似switch 监听channel读写
	concurrency.GoSelect()
	fmt.Printf("------------GoTimer------------\n")
	// GoTimer 定时器
	concurrency.GoTimer()
	fmt.Printf("------------GoTicker------------\n")
	// GoTicker 周期执行
	concurrency.GoTicker()
	fmt.Printf("------------GoAtomic------------\n")
	// Atomic包 原子变量 实现同步
	concurrency.GoAtomic()
}

func os() {
	fmt.Printf("------------GoFile------------\n")
	// 文件操作
	// goos.GoFile()
	// goos.GoFileWrite()
	// goos.GoFileRead()
	fmt.Printf("------------GoDir------------\n")
	// 文件夹操作
	// goos.GoDir()
	fmt.Printf("------------GoProcess------------\n")
	// 进程操作
	// goos.GoProcess()
	fmt.Printf("------------GoEnv------------\n")
	goos.GoEnv()
}
