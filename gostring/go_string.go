package gostring

import (
	"fmt"
	"strings"
)

func GoString() string {

	// 字符串可以 用 "" 或 `` 创建
	// "" 内可以包含转义字符 只能单行
	// 转义符
	// \r 回车
	// \n
	// \t
	// \'
	// \"
	// \\

	// `` 不支持转义 可以多行

	var str0 string = "hello world"
	var html string = `
		<html>
			<div>test</div>
		</html>
	`
	fmt.Printf("str1: %v\n", str0)
	fmt.Printf("html: %v\n", html)

	// 字符串拼接
	// golang中字符串都是不可变的 每次运算会产生一个新的字符串
	// 所以会产生很多临时无用的字符串 给gc带来压力

	str1 := "str1"
	str2 := "str2"
	str3 := str1 + "  " + str2
	fmt.Printf("str3: %v\n", str3)

	str4 := "str4"
	str4 += " "
	str4 += "str4"
	fmt.Printf("str4: %v\n", str4)

	// 使用fmt.Sprintf拼接字符串
	str5 := "str5"
	str6 := "str6"
	str7 := fmt.Sprintf("%s+%s", str5, str6)
	fmt.Printf("str7: %v\n", str7)

	// 使用string.Join()拼接字符串
	str8 := strings.Join([]string{str5, str6}, "+")
	fmt.Printf("str8: %v\n", str8)

	// 字符串切片
	str9 := "Hello test test!"
	n := 3
	m := 8
	// 获取字符串索引位置为 n 的原始字节
	fmt.Printf("str9[n]: %v\n", str9[n])
	// 获取 n 到 m-1 的字符串
	fmt.Printf("str9[n:m]: %v\n", str9[n:m])
	// 获取 n 到最后的字符串
	fmt.Printf("str9[n:]: %v\n", str9[n:])
	// 获取 开始到 m-1 的字符串
	fmt.Printf("str9[:m]: %v\n", str9[:m])

	// 字符串常用方法
	s := "hello&world!"
	// 长度
	fmt.Printf("len(s): %v\n", len(s))
	// 分割字符串 分割符是空
	fmt.Printf("strings.Split(s, \"\"): %v\n", strings.Split(s, ""))
	// 分割字符串
	fmt.Printf("strings.Split(s, \"w\"): %v\n", strings.Split(s, "&"))
	// 包含
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	// 前缀
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))
	// 后缀
	fmt.Printf("strings.HasSuffix(s, \"world!\"): %v\n", strings.HasSuffix(s, "world!"))
	// 第一个 l 下标
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))
	// 最后一个 l 下标
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))

	// 字符，字符串由字符组成
	// 字符有两种类型
	// unit8 byte 代表一个 ASCII码
	// rune 代表一个 UTF-8字符
	a := 'a'
	b := 'b'
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	return "learn-go-string"

}
