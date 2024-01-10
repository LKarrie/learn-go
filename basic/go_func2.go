package basic

import "fmt"

type Person3 struct {
	name string
}

func (p Person3) eat() {
	println(p.name + "eat")
}

func (p Person3) sleep() {
	println(p.name + "sleep")
}

func (p Person3) test1() {
	p.name = "change"
}

func (p *Person3) test2() {
	p.name = "change"
}

func GoFunc2() {

	// Go方法
	// Go中没有面向对象的特性 没有类对象的概念 但是可以使用结构体来模拟这些
	// 我们可以声明一些方法 属于某个结构体

	// Go中的方法是一种特殊的函数 定义于 struct 之上 与 struct 关联 绑定
	// 方法就是有接收者的函数

	// 普通函数
	// func test(test int) int {}
	// 方法
	// A struct {}
	// func (a A) test(test int) int {}

	p := Person3{"lk"}
	p.eat()
	p.sleep()

	// 注
	// 方法的 接收者类型 不一定非要是结构体 也可以是type定义的类型别名、slice、map、channel、func都可以
	// struct结合它的方法等价于面向对象中的类 struct可以和他的方法分开 并非要在一个文件 但是必须同包
	// (T type) 和 (T *type) 两种接收类型不同
	// 方法就是函数 不能重载 名称必须唯一
	// 如果 方法接收者(receiver)是一个指针类型 则会自动解除引用

	// Go方法接收者的类型

	// 方法的值类型和指针类型接收者 和 函数传递值类型和指向类型 相同
	p.test1()
	fmt.Printf("p: %v\n", p)
	// 传递了地址 p name属性改变
	p.test2()
	fmt.Printf("p: %v\n", p)

}
