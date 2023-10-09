package main

import (
	"fmt"
	"learn-go/goarray"
	"learn-go/goif"
	"learn-go/goloop"
	"learn-go/gooperator"
	"learn-go/goplaceholder"
	"learn-go/goslice"
	"learn-go/gostring"
	"learn-go/goswitch"
	"learn-go/govar"
	"learn-go/govartype"
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
	main2 := govartype.GoVarType()
	fmt.Printf("main2: %v\n", main2)

	fmt.Printf("------------GoString------------\n")
	//go中的字符字符串和常用方法
	main3 := gostring.GoString()
	fmt.Printf("main3: %v\n", main3)

	fmt.Printf("------------GoPlaceholder------------\n")
	//go中的占位符
	main4 := goplaceholder.GoPlaceholder()
	fmt.Printf("main4: %v\n", main4)

	fmt.Printf("------------GoOperator------------\n")
	//go中的运算符
	main5 := gooperator.GoOperator()
	fmt.Printf("main5: %v\n", main5)

	fmt.Printf("------------GoIf------------\n")
	//go中的if if中的goto
	main6 := goif.GoIf()
	fmt.Printf("main6: %v\n", main6)

	fmt.Printf("------------GoSwitch------------\n")
	//go中的switch switch中的break和fallthrough
	main7 := goswitch.GoSwitch()
	fmt.Printf("main7: %v\n", main7)

	fmt.Printf("------------GoLoop------------\n")
	//go中的for和range循环 循环中的break和goto
	main8 := goloop.GoLoop()
	fmt.Printf("main8: %v\n", main8)

	fmt.Printf("------------GoArray------------\n")
	//go中的数组 访问数组 和 遍历数组
	main9 := goarray.GoArray()
	fmt.Printf("main9: %v\n", main9)

	fmt.Printf("------------GoSlice------------\n")
	//go中的切片 切片的长度和容量 切片的初始化 空切片 切片的遍历
	main10 := goslice.GoSlice()
	fmt.Printf("main10: %v\n", main10)
}
