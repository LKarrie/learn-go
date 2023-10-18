package goextends

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	println("eat")
}

func (a Animal) sleep() {
	println("sleep")
}

type Dog struct {
	Animal
}

type Cat struct {
	a Animal
}

type Person struct {
	name string
	age  int
}

func (per Person) work() {
	println("work")
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name not null")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cant less than 0")
	}
	return &Person{name: name, age: age}, nil
}

func GoExtends() {
	// Go 本质没有继承的概念 可以使用 结构体嵌套实现
	dog := Dog{
		Animal{
			name: "dog",
			age:  2,
		},
	}
	cat := Cat{
		Animal{
			name: "cat",
			age:  2,
		},
	}
	dog.eat()
	cat.a.sleep()

	// 模拟类的 属性和方法
	per := Person{
		name: "lk",
		age:  18,
	}
	fmt.Printf("per: %v\n", per)
	per.work()

	// 模拟构造函数
	person, err := NewPerson("lk", 18)
	if err == nil {
		fmt.Printf("person: %v\n", person)
	}

}
