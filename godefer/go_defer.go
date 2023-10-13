package godefer

func GoDefer() string {

	// Go中的defer语句会将后面跟随的语句 进行延迟处理
	// 在 defer 归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
	// 也就是说 先被 defer 的语句将最后被执行 最后被 defer 的语句最先被执行

	// 关键字 defer 用于注册延迟调用
	// 这些调用 直到 return 前才被执行 因此可以用来做资源清理
	// 多个 defer 语句 按照先进后出的方式执行
	// defer 语句中的变量 在 defer 声明时就决定了

	// defer 用途
	// 关闭文件句柄
	// 锁资源释放
	// 数据库连接释放

	println("start")
	defer println("step1")
	defer println("step2")
	defer println("step3")
	println("end")

	return "learn-go-godefer"
}
