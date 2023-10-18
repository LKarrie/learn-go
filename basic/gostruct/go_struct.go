package gostruct

import "fmt"

// 也可以声明在顶端 函数可以访问
type Person struct {
	// id int
	// name string
	// age int
	// email string
	// 或者 合并为一行
	id, age     int
	name, email string
}

type A struct {
	id int
}
type B struct {
	a    A
	test string
}

func GoStruct() {

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

	// 结构体
	// Go 没有面向对象的概念 可以使用结构体实现面向对象的特性
	// type Person struct {
	// 	// id int
	// 	// name string
	// 	// age int
	// 	// email string
	// 	// 或者 合并为一行
	// 	id, age     int
	// 	name, email string
	// }

	var lk Person
	fmt.Printf("lk: %v\n", lk)

	// 访问结构体成员 使用 . 访问对应属性
	lk.id = 1
	lk.name = "lkarrie"
	lk.age = 18
	lk.email = "test@123.com"
	fmt.Printf("lk: %v\n", lk)

	// 匿名结构体
	var a struct {
		id1, id2 int
	}
	a.id1 = 1
	a.id2 = 2

	// 结构体初始化
	lk2 := Person{
		id:   1,
		name: "lk",
		age:  18,
		// 可以部分初始化
		// email: "123",
	}
	fmt.Printf("lk2: %v\n", lk2)

	// 这种方式初始化 必须初始化结构体 所有字段 顺序需一致
	lk3 := Person{
		1, 18, "123", "123",
	}
	fmt.Printf("lk3: %v\n", lk3)

	// 结构体指针
	var p *Person = &lk2
	fmt.Printf("p: %p\n", p)
	fmt.Printf("*p: %v\n", *p)

	// new 创建结构体指针
	var p2 = new(Person)
	fmt.Printf("p2: %T\n", p2)
	// 访问 结构体指针成员
	p2.age = 1
	p2.name = "test"
	fmt.Printf("p2: %v\n", p2)

	// 结构体做函数参数
	// 直接传递
	test1(lk)
	// 传递结构体指针
	test2(p2)

	// 嵌套结构体 实现继承的效果
	// B 包含 A
	var b B
	b.test = "test"
	b.a.id = 1
	fmt.Printf("b: %v\n", b)

}

func test1(p Person) {
	p.id = 1000
	fmt.Printf("p: %v\n", p)
}

func test2(p *Person) {
	p.id = 1000
	fmt.Printf("p: %v\n", p)
}
