package gomap

import "fmt"

func GoMap() string {

	// map是引用类型 内部实现是哈希表

	// map声明
	var m1 map[string]string = make(map[string]string)
	m1["key1"] = "value1"

	fmt.Printf("m1: %v\n", m1)

	m2 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	fmt.Printf("m2: %v\n", m2)

	// map访问
	value1 := m2["key1"]
	value2 := m2["key2"]
	fmt.Printf("value1: %v\n", value1)
	fmt.Printf("value2: %v\n", value2)

	// 判断某个键是否存在
	// value,ok := map[key]
	v, ok := m2["key3"]
	if ok {
		fmt.Printf("key3存在\n")
		fmt.Printf("v: %v\n", v)
	} else {
		fmt.Printf("key3不存在\n")
	}

	// 遍历map
	// 遍历key
	for key := range m2 {
		fmt.Printf("key: %v\n", key)
	}
	// 遍历key和value
	for key, value := range m2 {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("value: %v\n", value)
	}

	return "learn-go-gomap"
}
