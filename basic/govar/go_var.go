package govar

import "fmt"

func getTwoNames() (string, string) {
	return "param1", "param2"
}

func GoVar() {

	// 标识符只能数字 字母 下划线 组成
	// 标识符只能字母下划线开头
	// 标识符区分大小写

	// 单个声明
	var name0 string
	// 批量声明
	var (
		name1 string
		name2 string
	)
	// 初始化
	// 整数浮点 初始化后默认 0
	// 字符串 初始化后默认 “”
	// 布尔 初始化后默认 false
	// 切片、函数、指针变量类型默认 nil

	// 批量初始化
	var name3, name4 = "name3", "name4"

	// 短变量声明
	name5 := "name5"

	// 匿名变量
	name6, _ := getTwoNames()

	// 常量
	const const0 = "const0"
	const (
		const1 = "const1"
		const2 = "const2"
	)

	// 常量 不可改变
	// const0 = "const0"

	// iota
	// 自增常量 从0开始
	const (
		iota0 = iota
		_     = iota
		iota2 = "插队"
		iota3 = iota
	)

	fmt.Printf("name0: %v\n", name0)
	fmt.Printf("name1: %v\n", name1)
	fmt.Printf("name2: %v\n", name2)
	fmt.Printf("name3: %v\n", name3)
	fmt.Printf("name4: %v\n", name4)
	fmt.Printf("name5: %v\n", name5)
	fmt.Printf("name6: %v\n", name6)

	fmt.Printf("const0: %v\n", const0)
	fmt.Printf("const1: %v\n", const1)
	fmt.Printf("const2: %v\n", const2)
	fmt.Printf("iota0: %v\n", iota0)
	fmt.Printf("iota2: %v\n", iota2)
	fmt.Printf("iota3: %v\n", iota3)

}
