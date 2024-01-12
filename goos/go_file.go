package goos

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("./goos/test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
	// 很重要
	f.Close()
}

// 重命名文件
func renameFile() {
	err := os.Rename("./goos/test.txt", "./goos/test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 写文件
func writeFile() {
	s := "hello world"
	os.WriteFile("./goos/test2.txt", []byte(s), os.ModePerm)
}

// 读文件
func readFile() {
	b, err := os.ReadFile("./goos/test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b[:]))
	}
}

func GoFile() {
	createFile()
	renameFile()
	writeFile()
	readFile()
}
