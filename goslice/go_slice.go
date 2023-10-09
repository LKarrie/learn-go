package goslice

import "fmt"

func GoSlice() string {

	// 切片就是可变长的数组 底层使用数组实现 增加自动扩容功能
	var slice1 []string
	var slice2 []int
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)

	fmt.Printf("(slice1 == nil): %v\n", (slice1 == nil))
	fmt.Printf("(slice2 == nil): %v\n", (slice2 == nil))

	// 切片的长度和容量
	// 切片的长度是它所包含的元素个数
	// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数
	// 切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取

	var slice3 = []string{"a", "b", "c", "d"}
	var slice4 = []int{1, 2, 3}
	fmt.Printf("len(slice3): %v\n", len(slice3))
	fmt.Printf("cap(slice3): %v\n", cap(slice3))

	fmt.Printf("len(slice4): %v\n", len(slice4))
	fmt.Printf("cap(slice4): %v\n", cap(slice4))

	var slice5 = make([]bool, 2, 5)
	fmt.Printf("len(slice5): %v\n", len(slice5))
	fmt.Printf("cap(slice5): %v\n", cap(slice5))
	fmt.Printf("slice5: %v\n", slice5)

	//默认情况下切片cap值小于1024的时候是成倍数增长的
	//比如有1变为2在变为4再变为16……但是超过1024后就变为原切片的1.25倍
	slice5 = append(slice5, true, true, true, true)
	fmt.Printf("len(slice5): %v\n", len(slice5))
	fmt.Printf("cap(slice5): %v\n", cap(slice5))
	fmt.Printf("slice5: %v\n", slice5)

	// 详细理解 切片的长度和容量
	fmt.Printf("-------------详细理解 切片的长度和容量-------------\n")
	s := []int{2, 3, 5, 7, 11, 13}
	// 第一个输出为[2,3,5,7,11,13]，长度为6，容量为6
	printSlice(s)

	// 截取切片使其长度为 0
	// 第一个输出为[2,3,5,7,11,13]，长度为6，容量为6
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	// 左指针指向s[0]，右指针指向s[4]，由于切片概念是只包含左边元素不包含右边元素，所以长度为4，但左指针在s[0]处，走过0个元素，所以容量仍然为6
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	// 在经历步骤3切片后的基础上，左指针指向s[2]，右指针指向最右边，所以长度为2，由于左指针走过两个元素，离最右边还剩4个元素，所以容量为4
	s = s[2:]
	printSlice(s)

	// 输出
	// len=6 cap=6 [2 3 5 7 11 13]
	// len=0 cap=6 []
	// len=4 cap=6 [2 3 5 7]
	// len=2 cap=4 [5 7]
	fmt.Printf("--------------------------\n")

	// 切片的切分
	// 使用数组的部分元素初始化（切片表达式）
	var s1 = []int{1, 2, 3, 4, 5}
	s2 := s1[0:3]
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[2:5]
	fmt.Printf("s4: %v\n", s4)
	s5 := s1[:]
	fmt.Printf("s5: %v\n", s5)

	// 切片的初始化
	// 使用数组初始化
	arr := [...]int{1, 2, 3}
	s6 := arr[:]
	fmt.Printf("s6: %v\n", s6)

	// 空（nil）切片
	// 一个切片在未初始化之前默认未nil 长度0 容量0
	var s7 []int
	fmt.Printf("(s7 == nil): %v\n", (s7 == nil))
	fmt.Printf("len(s7): %v\n", len(s7))
	fmt.Printf("cap(s7): %v\n", cap(s7))

	// 切片遍历
	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[%d]: %v\n", i, s1[i])
	}
	for i, v := range s1 {
		fmt.Printf("s1[%d]: %v\n", i, v)
	}

	// 切片的添加和删除
	// 添加
	s8 := []int{}
	s8 = append(s8, 1)
	s8 = append(s8, 2)
	s8 = append(s8, 3, 4, 5)
	fmt.Printf("s8: %v\n", s8)
	// 添加另外一个切片
	s8 = append(s8, s1...)
	fmt.Printf("s8: %v\n", s8)

	// 删除
	// 要从切片s9中删除索引为index的元素 a = append(a[:index],a[index+1]...)
	s9 := []int{1, 2, 3, 4, 5}
	s9 = append(s9[:2], s9[3:]...)
	fmt.Printf("s9: %v\n", s9)

	// 拷贝
	s10 := []int{1, 2, 3}
	s11 := s10
	// s10 调整 s11 跟着变化
	s10[0] = 100
	fmt.Printf("s10: %v\n", s10)
	fmt.Printf("s11: %v\n", s11)
	s12 := make([]int, 3)
	copy(s12, s10)
	// s12 不变化
	s10[0] = 4
	fmt.Printf("s10: %v\n", s10)
	fmt.Printf("s12: %v\n", s12)

	return "learn-go-slice"
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
