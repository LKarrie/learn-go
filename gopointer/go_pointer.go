package gopointer

import "fmt"

func GoPointer() string {

	// Go语言中传递参数是值拷贝，当我们想要修改某个变量时候，可以创建一个指向该变量地址的指针变量，传递数据使用指针无需拷贝数据
	// 指针类型不能进行 偏移和运算
	// Go语言中指针操作：&（取地址）和 *（根据地址取值）

	// 指针地址和指针类型
	// 每个变量都在运行时拥有一个地址，地址代表变量在内存中的位置，Go使用 & 放在变量前对变量进行取地址操作
	// Go 语言中的值类型 都有对应的执行类型 例如 *int *int64 *string

	// 指针语法
	// 一个指针变量指向了一个值内存地址（声明一个指针之后，可以像变量赋值一样，把一个值的内存地址存放到指针中）
	// 类似变量和常亮 在指针前需要声明指针
	// var var_name *var_type

	// 声明变量
	var a int = 20
	// 声明指针变量
	var ip *int
	// 指针变量的存储地址
	ip = &a

	println(&a)
	fmt.Printf("ip变量存储的指针地址: %v\n", ip)
	fmt.Printf("ip变量的值: %v\n", *ip)

	// 指向数组的指针
	b := []int{1, 3, 5}
	var ptr [3]*int

	for i := 0; i < 3; i++ {
		// 取地址
		ptr[i] = &b[i]
	}

	for i := 0; i < 3; i++ {
		// 取值
		fmt.Printf("b[%d] = %d\n", i, *ptr[i])
	}

	return "learn-go-pointer"
}
