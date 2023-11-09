package main

import (
	"fmt"
	"learn-go/basic/goarray"
	"learn-go/basic/godefer"
	"learn-go/basic/goextends"
	"learn-go/basic/gofunc"
	"learn-go/basic/gofunc2"
	"learn-go/basic/goif"
	"learn-go/basic/goinit"
	"learn-go/basic/gointerface"
	"learn-go/basic/goloop"
	"learn-go/basic/gomap"
	"learn-go/basic/goocp"
	"learn-go/basic/gooperator"
	"learn-go/basic/goplaceholder"
	"learn-go/basic/gopointer"
	"learn-go/basic/goslice"
	"learn-go/basic/gostring"
	"learn-go/basic/gostruct"
	"learn-go/basic/goswitch"
	"learn-go/basic/govar"
	"learn-go/basic/govartype"
	"learn-go/concurrency/gochannel"
	"learn-go/concurrency/gogoroutines"
	"learn-go/concurrency/gomutex"
	"learn-go/concurrency/goruntime"
	"learn-go/concurrency/goselect"
	"learn-go/concurrency/goticker"
	"learn-go/concurrency/gotimer"
	"learn-go/concurrency/gowaitgroup"
)

func main() {
	// 基础
	// gobasic()
	// 并发编程
	goconcurrency()
}

func gobasic() {
	fmt.Printf("------------GoVar------------\n")
	//go中的变量声明和初始化
	//go中的常量
	//go中得iota
	govar.GoVar()

	fmt.Printf("------------GoVarType------------\n")
	//go中的整数浮点数复数
	govartype.GoVarType()

	fmt.Printf("------------GoString------------\n")
	//go中的字符字符串和常用方法
	gostring.GoString()

	fmt.Printf("------------GoPlaceholder------------\n")
	//go中的占位符
	goplaceholder.GoPlaceholder()

	fmt.Printf("------------GoOperator------------\n")
	//go中的运算符
	gooperator.GoOperator()

	fmt.Printf("------------GoIf------------\n")
	//go中的if if中的goto
	goif.GoIf()

	fmt.Printf("------------GoSwitch------------\n")
	//go中的switch switch中的break和fallthrough
	goswitch.GoSwitch()

	fmt.Printf("------------GoLoop------------\n")
	//go中的for和range循环 循环中的break和goto 递归
	goloop.GoLoop()

	fmt.Printf("------------GoArray------------\n")
	//go中的数组 访问数组 和 遍历数组
	goarray.GoArray()

	fmt.Printf("------------GoSlice------------\n")
	//go中的切片 切片的长度和容量 切片的初始化 空切片 切片的遍历
	goslice.GoSlice()

	fmt.Printf("------------GoMap------------\n")
	//go中的Map 判断Map键是否存在 遍历Map的key和value
	gomap.GoMap()

	fmt.Printf("------------GoFunc------------\n")
	//go中的函数 函数类型 函数的返回值 函数的参数 函数类型和函数变量(函数作为参数) 函数做返回值 匿名函数 闭包
	gofunc.GoFunc()

	fmt.Printf("------------GoDefer------------\n")
	//go中的defer语句
	godefer.GoDefer()

	fmt.Printf("------------GoInit------------\n")
	//go中的init语句
	goinit.GoInit()

	fmt.Printf("------------GoPointer------------\n")
	//go中的指针 指向数组的指针
	gopointer.GoPointer()

	fmt.Printf("------------GoStruct------------\n")
	//go中的结构体 结构体访问 初始化 结构体指针 结构体做参 嵌套结构体（继承）
	gostruct.GoStruct()

	fmt.Printf("------------GoFunc2------------\n")
	//go中的方法 方法的接收者类型
	gofunc2.GoFunc2()

	fmt.Printf("------------GoInterface------------\n")
	//go中的接口 接口实现 接口接收者类型 接口嵌套
	gointerface.GoInterface()

	fmt.Printf("------------GoOcp------------\n")
	//go OCP 设计原则（面向对象可复用 开-闭原则 面对扩展开放 面对修改关闭）
	goocp.GoOcp()

	fmt.Printf("------------GoExtends------------\n")
	//go 继承 模拟类属性和方法 模拟类构造函数
	goextends.GoExtends()
}

func goconcurrency() {
	fmt.Printf("------------GoGoroutines------------\n")
	// 协程
	gogoroutines.GoGoroutines()
	fmt.Printf("------------GoChannel------------\n")
	// channel通道 channel遍历
	gochannel.GoChannel()
	fmt.Printf("------------GoWaitGroup------------\n")
	// WaitGroup同步
	gowaitgroup.GoWaitGroup()
	fmt.Printf("------------GoRuntime------------\n")
	// runtime包中的一些方法
	goruntime.GoRuntime()
	fmt.Printf("------------GoMutex------------\n")
	// mutex包中lock实现同步
	gomutex.GoMutex()
	fmt.Printf("------------GoSelect------------\n")
	// GoSelect 类似switch 监听channel读写
	goselect.GoSelect()
	fmt.Printf("------------GoTimer------------\n")
	// GoTimer 定时器
	gotimer.GoTimer()
	fmt.Printf("------------GoTicker------------\n")
	// GoTicker 周期执行
	goticker.GoTicker()
}
