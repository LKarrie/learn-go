package main

import (
	"fmt"
	"learn-go/gooperator"
	"learn-go/goplaceholder"
	"learn-go/gostring"
	"learn-go/govar"
	gotvarype "learn-go/govartype"
)

func main() {

	fmt.Printf("------------GoVar------------\n")
	//go中的变量声明和初始化
	//go中的常量
	//go中得iota
	main := govar.GoVar()
	fmt.Printf("main: %v\n", main)

	fmt.Printf("------------GoVarType------------\n")
	//go中的整数浮点数复数
	main2 := gotvarype.GoVarType()
	fmt.Printf("main2: %v\n", main2)

	fmt.Printf("------------GoString------------\n")
	//go中的字符字符串和常用方法
	main3 := gostring.GoString()
	fmt.Printf("main3: %v\n", main3)

	fmt.Printf("------------GoPlaceholder------------\n")
	//go中的占位符
	main4 := goplaceholder.GoPlaceholder()
	fmt.Printf("main4: %v\n", main4)

	fmt.Printf("------------GoPlaceholder------------\n")
	//go中的运算符
	main5 := gooperator.GoOperator()
	fmt.Printf("main5: %v\n", main5)

}
