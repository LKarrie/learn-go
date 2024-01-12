package goos

import (
	"fmt"
	"os"
)

// 如果文件路径指定的文件不存在 会输出报错
// The system cannot find the file specified.

func write() {
	f, err := os.OpenFile("./goos/a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte("hello golang"))
	f.Close()
}

func writeString() {
	f, err := os.OpenFile("./goos/a.txt", os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("hello java")
	f.Close()
}

func writeAt() {
	f, err := os.OpenFile("./goos/a.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
	}
	// 第三个字符开始 替换成 aaa
	f.WriteAt([]byte("aaa"), 3)
	f.Close()
}

func GoFileWrite() {
	f0, err := os.Create("./goos/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	f0.Close()
	write()
	writeString()
	writeAt()
}
