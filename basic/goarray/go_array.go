package goarray

import "fmt"

func GoArray() {
	// 数组的定义
	// var var_name [SIZE] var_type

	var a [3]int
	var b [2]string

	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)

	// 没有初始化的数组 默认元素都是零值 布尔是false 字符串是空字符串

	var c [3]int
	var d [2]string
	var e [2]bool

	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

	// 初始化列表
	var aa = [3]int{1, 2, 3}
	var bb = [2]string{"a", "b"}
	var cc = [2]bool{true, false}

	fmt.Printf("aa: %v\n", aa)
	fmt.Printf("bb: %v\n", bb)
	fmt.Printf("cc: %v\n", cc)

	// 省略数组长度
	var aaa = [...]int{1, 2, 3}
	var bbb = [...]string{"a", "b"}
	var ccc = [...]bool{true, false}
	fmt.Printf("aaa: %v\n", aaa)
	fmt.Printf("bbb: %v\n", bbb)
	fmt.Printf("ccc: %v\n", ccc)

	// 按索引初始化
	var aaaa = [...]int{0: 1, 2: 3}
	fmt.Printf("aaaa: %v\n", aaaa)

	// 访问数组
	var test [2]int
	// 赋值
	test[0] = 100
	test[1] = 200
	fmt.Printf("test[0]: %v\n", test[0])
	fmt.Printf("test[1]: %v\n", test[1])
	// 修改
	test[0] = 1
	test[1] = 2
	fmt.Printf("test[0]: %v\n", test[0])
	fmt.Printf("test[1]: %v\n", test[1])

	// 遍历数组
	test1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(test1); i++ {
		fmt.Printf("test1[%d]: %v\n", i, test1[i])
	}

	for i, v := range test1 {
		fmt.Printf("test1[%d]: %v\n", i, v)
	}

}
