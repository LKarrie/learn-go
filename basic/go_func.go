package basic

import "fmt"

// 函数类型
type fun func(int, int) int

func GoFunc() {

	// 函数包含 函数的名称 参数列表 和 返回值类型 这些构成函数的签名
	// go 语言中函数的特性
	// go 语言中有三种函数：普通函数、匿名函数（无名称的函数）、方法（定义在struct上的函数）
	// go 语言中不允许函数重载，不允许函数同名
	// go 语言中的函数不能嵌套函数，但是可以嵌套匿名函数
	// 函数是一个值，可以将函数复制给变量，使这个变量也成为函数
	// 函数可以作为参数传递给另外一个函数
	// 函数的返回值可以是一个函数
	// 函数调用的时候，如果有参数传递给函数，则先拷贝参数的副本，再将副本传递给函数
	// 函数参数可以没有名称

	// 函数的定义
	// func：函数由func开始声明
	// func_name：函数名称，函数名和参数列表构成参数签名
	// [param list]：参数列表，参数列表指定参数类型、顺序、参数个数
	// [return_types]：返回类型
	// 函数体：函数定义的代码集合

	// func func_name( [param list] ) [return_types] {
	// 		函数体
	// }

	// 调用函数
	s := sum(1, 2)
	fmt.Printf("s: %v\n", s)

	max := compare(1, 2)
	fmt.Printf("max: %v\n", max)

	// 函数返回值
	// 函数可以没有返回值
	// 函数可以有 一个 或者 多个返回值
	// 返回值名称没有使用的两个情况
	f1()
	f2()
	// Go 经常会使用一个返回值作为函数是否执行成功、是否有错误信息的判断条件
	// 例如：return value,exists、return value,ok，return value,err等

	// 可以使用 _ 丢弃返回值
	_, x := f3()
	fmt.Printf("x: %v\n", x)

	// 函数参数
	// 函数可以有 0个 或者 多个参数，参数需要指定类型
	// 声明函数时的参数列表叫做形参，调用时传递的参数叫做实参
	// Go 中函数参数传参 是传值的方式
	// Go 中参数可以使用变长参数

	// 值传递
	var a int = 100
	f4(a)
	fmt.Printf("a: %v\n", a)
	// map slice interface channel 这些数据类型本身就是指针类型 就算拷贝传值也是拷贝指针
	b := []int{1, 2}
	fmt.Printf("b: %v\n", b)
	// 调整指针类型会影响旧值
	f5(b)
	fmt.Printf("b: %v\n", b)

	// 变长参数
	f6("name", 1, 2, 3, 4, 5, 6, 7, 8)

	// 函数类型赋值
	var f fun = sum
	i := f(1, 2)
	fmt.Printf("i: %v\n", i)

	// 函数做参数
	f8("hello", f7)

	// 函数做返回值
	f9 := cal("+")
	i2 := f9(1, 2)
	fmt.Printf("i2: %v\n", i2)

	// 匿名函数
	// Go函数不能嵌套 函数内部可以定义匿名函数
	inner := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	i3 := inner(1, 2)
	fmt.Printf("i3: %v\n", i3)

	// 匿名函数自己执行
	func(a int, b int) {
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Printf("max: %v\n", max)
	}(1, 2)

	// Go闭包
	// 闭包可以理解成 定义在一个函数内部的函数 在本质上 闭包是将函数内部和函数外部连接起来的桥梁
	// 或者说是 函数和其引用环境的组合

	// 闭包指的是一个函数和与其应用的环境组合而成的实体 简单来说 闭包 = 函数 + 引用环境
	var ff = add()
	// ff 内部 形成闭包 x 第一次是 0
	fmt.Println(ff(10))
	// x 第二次是 10
	fmt.Println(ff(20))
	// x 第三次是 30
	fmt.Println(ff(30))
	fmt.Println("-----------")

	// f1 形成新的闭包 第一次还是 0
	f1 := add()
	fmt.Println(f1(40))
	// x 第二次 40
	fmt.Println(f1(50))

}

// 返回一个函数
func add() func(int) int {
	// x 不初始化默认是0
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 定义一个函数
// 求和
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

// 比较大小
func compare(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func f1() (name string, age int) {
	name = "test"
	age = 18
	// 等价于 return name, age
	return
}

func f2() (name string, age int) {
	n := "test"
	a := 18
	// return 覆盖命名返回值，返回值名称没有被使用
	return n, a
}

func f3() (int, int) {
	return 1, 2
}

func f4(a int) {
	a = 200
	fmt.Printf("a: %v\n", a)
}

func f5(b []int) {
	b[0] = 100
}

func f6(name string, args ...int) {
	fmt.Printf("name: %v\n", name)
	for v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func f7(name string) {
	fmt.Printf("name: %v\n", name)
}

func f8(name string, f func(string)) {
	f(name)
}

func cal(s string) func(int, int) int {
	switch s {
	case "+":
		return sum
	default:
		return nil
	}
}
