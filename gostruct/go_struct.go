package gostruct

import "fmt"

func GoStruct() string {

	// Go语言定义类型
	// type newtype Type
	type MyInt int
	// i 为MyInt类型
	var i MyInt = 100
	fmt.Printf("i: %v\n", i)

	// Go语言类型别名
	// type newType = Type
	// 类型别名定义
	type MyInt2 = int
	// i 其实还是 int
	var i2 MyInt2 = 100
	fmt.Printf("i2: %v\n", i2)

	// 类型定义和类型别名的区别
	// 类型定义相当于 定义了一个全新的类型 与之前的类型不同，别名只是替换一个名字
	// 类别名编译后会消失
	// 类别名和原来的类型一致 所以有以前类有的方法，重新定义一个新的类型则没有

	return "learn-go-struct"
}
