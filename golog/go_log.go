package golog

import (
	"fmt"
	"log"
	"os"
)

func logTypes() {
	// golang 内置的 log 包 实现简单日志服务
	// 调用 log 包函数 可实现简单的日志打印

	// log 包中钟类型的日志函数
	// print	只打印
	// panic	打印 并 抛出 panic 异常
	// fatal	打印 并 强制结束程序 os.Exit(1) defer 函数不会执行

	defer fmt.Println("发生了 panic错误")
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20
	log.Println(name, ",", age)
	// exit status 2
	log.Panic("致命错误1")

	// exit status 1
	// log.Fatal("致命错误2")
	fmt.Println("end...")
}

const ()

func logFormat() {
	// 默认情况 log 只会打印时间
	// log 包提供定制的接口 打印其他信息
	// 返回标准log输出配置
	// func Flags() int
	// 设置标准log输出配置
	// func SetFlags(flag int)

	i := log.Flags()
	fmt.Printf("i: %v\n", i)

	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	// Ldate         = 1 << iota     // 日期：2009/01/23
	// Ltime                         // 时间：01:23:23
	// Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
	// Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	// Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	// LUTC                          // 使用UTC时间
	// 默认 值为3
	// LstdFlags     = Ldate | Ltime // 标准logger的初始值

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("my log")

	// 日志前缀
	// 返回日志的前缀配置
	// func Prefix() string
	// 设置日志前缀
	// func SetPrefix(prefix string)

	s := log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.SetPrefix("MyLog: ")
	s = log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.Print("my log...")
}

func logToFile() {
	f, err := os.OpenFile("my.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	log.SetOutput(f)
	log.Print("my log...")
}

func logPosition() {
	f, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	log.SetOutput(f)
	log.Print("my log...")
}

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger = log.New(logFile, "success", log.Ldate|log.Ltime|log.Lshortfile)
}

func GoLog() {
	// 日志几种类型
	// logTypes()
	// 日志格式
	// logFormat()
	// 日志写文件
	// logToFile()
	// 日志输出位置

	// logger.Println("自定义logger")
}
