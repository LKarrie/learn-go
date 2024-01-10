package basic

type Pet interface {
	eat()
	sleep()
}

type Dog1 struct {
}

func (d Dog1) eat() {
	println("Dog1 eat")
}

func (d Dog1) sleep() {
	println("Dog1 sleep")
}

type Cat1 struct {
}

func (d Cat1) eat() {
	println("Cat1 eat")
}

func (d Cat1) sleep() {
	println("Cat1 sleep")
}

type Person2 struct {
	name string
}

func (person2 Person2) have(p Pet) {
	p.eat()
	p.sleep()
}

func GoOcp() {

	// 面向对象可复用设计 开-闭 原则
	// 对扩展开放 对修改关闭
	// 解耦代码

	person := Person2{"lk"}
	d := Dog1{}
	c := Cat1{}
	person.have(d)
	person.have(c)

}
