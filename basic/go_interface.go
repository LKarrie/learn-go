package basic

type USB interface {
	read()
	write()
}

type Computer struct {
}

type Mobile struct {
}

func (c Computer) read() {
	println("c read")
}
func (c Computer) write() {
	println("c write")
}
func (m Mobile) read() {
	println("m read")
}
func (m Mobile) write() {
	println("m write")
}

// type test interface {
// 	test1()
// 	test2()
// }

// type A struct {
// }

// func (a A) test1() {
// 	println("test1")
// }

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

// 嵌套接口
type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
}

func (f Fish) fly() {
	println("fly")
}
func (f Fish) swim() {
	println("swim")
}

func GoInterface() {

	// Go接口会定义通用规范
	// Go接口是一种新的 类型定义 他把具有共性的方法定义在一起 任何其他类型只要实现了这些方法就是 实现了这个接口

	c := Computer{}
	m := Mobile{}
	c.read()
	c.write()
	m.read()
	m.write()

	// 实现接口 必须实现接口中的 所有方法

	// 报错 A does not implement test(missing method test2)
	// var t test = A{}

	// Go接口值类型接收者和指针类型接收者
	// 和 传参 相同 值类型接收者就是拷贝 是一个副本 指针类型接收者是传递指针

	// Go接口和类型的关系
	// 一个类型可以实现多个接口
	// 多个类型可以实现同一个接口（多态）

	// 接口也可以嵌套
	var ff FlyFish = Fish{}
	ff.fly()
	ff.swim()
}
