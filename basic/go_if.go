package basic

import (
	"fmt"
)

func GoIf() {
	// if 语句
	// 不需要使用括号将条件包起来
	// 大括号 {} 必须存在 即使只有一行一句
	// 左括号必须在 if 或 else 的同一行
	// 在 if 之后 条件语句之前 可以添加变量初始化的语句 使用 ; 分隔
	// 不能使用 0 非0 表示真假
	if num0 := 1; num0 > 0 {
		fmt.Printf("\"if : \": %v\n", "if")
	} else {
		fmt.Printf("\"else : \": %v\n", "else")
	}

	num1 := 3
	if num1 >= 3 {
		fmt.Printf("\">3\": %v\n", ">3")
	} else if num1 > 1 && num1 < 3 {
		fmt.Printf("\"1-3\": %v\n", "1-3")
	} else {
		fmt.Printf("\"<1\": %v\n", "<1")
	}

	// goto 可以跳出 if
	num2 := 1
	if num2 == 1 {
		goto LABEL1
	} else {
		fmt.Printf("other\n")
	}

LABEL1:
	fmt.Printf("next\n")

}
