package goinit

func GoInit() string {

	// Go 中的特殊函数 init 函数，先于main函数执行，实现包级别的一些初始化操作

	// init函数先于main函数自动执行，不能被其他函数调用
	// init函数没有输入参数、返回值
	// 每个包可以有多个init函数
	// 包的每个资源文件也可以有多个init函数
	// 同一个包的init执行顺序 golang没有明确定义 不要依赖这个执行顺序进行编程
	// 不同包的init函数按照包导入的依赖关系决定执行顺序

	// Golang 初始化顺序
	// 变量初始化 -> init() -> main()

	return "learn-go-init"
}

func init() {
	println("init2")
}

func init() {
	println("init")
}

func initVar() int {
	println("int var...")
	return 100
}
