package goocp

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
}

func (d Dog) eat() {
	println("Dog eat")
}

func (d Dog) sleep() {
	println("Dog sleep")
}

type Cat struct {
}

func (d Cat) eat() {
	println("Cat eat")
}

func (d Cat) sleep() {
	println("Cat sleep")
}

type Person struct {
	name string
}

func (person Person) have(p Pet) {
	p.eat()
	p.sleep()
}

func GoOcp() {

	// 面向对象可复用设计 开-闭 原则
	// 对扩展开放 对修改关闭
	// 解耦代码

	person := Person{"lk"}
	d := Dog{}
	c := Cat{}
	person.have(d)
	person.have(c)

}
