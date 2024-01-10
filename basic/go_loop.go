package basic

import (
	"fmt"
)

func GoLoop() {

	//for循环
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %v\n", i)
	}

	//for循环初始条件可以在外部
	i := 1
	for ; i <= 3; i++ {
		fmt.Printf("i: %v\n", i)
	}

	//初始和结束条件都可以省略 类似while
	for i <= 5 {
		fmt.Printf("i: %v\n", i)
		i++
	}
	fmt.Printf("i: %v\n", i)

	//永远循环
	// for {

	// }

	// for循环可以通过break goto return panic 语句退出循环

	//range循环
	//循环数组
	var a = [3]int{1, 2, 3}
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
	//循环字符串
	var b = "我是菜鸟"
	for i, v := range b {
		fmt.Printf("i: %d, v: %c\n", i, v)
	}
	//循环切片
	var s = []int{1, 2, 3}
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
	//循环map
	m := make(map[string]string)
	m["name1"] = "noob1"
	m["name2"] = "noob2"
	//k是key v是value
	for k, v := range m {
		fmt.Printf("k: %v,v: %v\n", k, v)
	}

	//循环中break可以跳转到标签
MY_LABEL:
	for i := 0; i < 4; i++ {
		if i == 2 {
			// 到标签处了 不在执行for循环
			break MY_LABEL
			// break
		}
		fmt.Printf("i: %v\n", i)
	}

	// 循环中的goto
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				goto LABEL1
			}
		}
	}
LABEL1:
	fmt.Println("label1")

	// 递归
	n := 5
	r := aa(n)
	fmt.Printf("r: %v\n", r)

}

func aa(n int) int {
	// 返回条件
	if n == 1 {
		return 1
	} else {
		// 自己调用自己
		return n * aa(n-1)
	}
}
