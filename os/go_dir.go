package os

import (
	"fmt"
	"os"
)

// 创建目录
func createDir() {
	// 创建单个目录
	err := os.Mkdir("./os/test", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err = os.MkdirAll("./os/test/a/b", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除目录 会删除目录下的所有目录和内容
func removeDir() {
	err := os.RemoveAll("./os/test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 获得工作目录
func getWd() {
	// 这个工作目录是 当前main.go的目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

// 修改工作目录
func chWd() {
	err := os.Chdir("d:/")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
}

// 获得临时目录
func getTemp() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func GoDir() {
	createDir()
	removeDir()
	getWd()
	chWd()
	getTemp()
}
