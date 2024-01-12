package goos

import (
	"fmt"
	"os"
)

// 打开关闭文件
func openCloseFile() {
	// 只能读
	f, err := os.Open("./goos/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("f.Name(): %v\n", f.Name())
	// 根据第二个参数 可以读写或者创建
	// 二次创建会覆盖源文件
	f2, _ := os.OpenFile("./goos/a1.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f2.Name(): %v\n", f2.Name())

	err = f.Close()
	fmt.Printf("err: %v\n", err)
	err2 := f2.Close()
	fmt.Printf("err2: %v\n", err2)
}

// 创建文件
func createFile2() {
	// 等价于：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	f, _ := os.Create("./goos/a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	// 第一个参数 目录默认：Temp 第二个参数 文件名前缀
	// tmp 文件名类似 temp1211373414
	f2, _ := os.CreateTemp("./goos", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

// 读操作
func readOps() {
	// 循环读取
	/* 	f, _ := os.Open("./goos/a.txt")
	   	for {
	   		buf := make([]byte, 6)
	   		n, err := f.Read(buf)
	   		fmt.Println(string(buf))
	   		fmt.Printf("n: %v\n", n)
	   		if err != nil {
	   			fmt.Println(err)
	   		}
	   		if err == io.EOF {
	   			break
	   		}
	   	}
	   	f.Close() */

	/* 	buf := make([]byte, 10)
	   	f2, err2 := os.Open("./goos/a.txt")
	   	if err2 != nil {
	   		fmt.Println(err2)
	   	}
	   	// 从5开始读10个字节
	   	n0, _ := f2.ReadAt(buf, 5)
	   	fmt.Printf("n0: %v\n", n0)
	   	fmt.Printf("string(buf): %v\n", string(buf))
	   	f2.Close() */

	// 读取目录下的所有内容 包括目录 只会读一层 不会继续读下层目录
	/* 	f, _ := os.Open("goos")
	   	de, _ := f.ReadDir(-1)
	   	for _, v := range de {
	   		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
	   		fmt.Printf("v.Name(): %v\n", v.Name())
	   	} */

	// 定位
	f3, _ := os.Open("./goos/a.txt")
	// io.SeekStart 0
	// io.SeekCurrent 1
	// io.SeekEnd 2

	// 1234567 Seek(3,0) 输出 4567
	ret, err := f3.Seek(3, 0)
	fmt.Printf("ret: %v\n", ret)
	fmt.Printf("err: %v\n", err)
	buf := make([]byte, 10)
	n, _ := f3.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f3.Close()

}

func GoFileRead() {
	// 二次创建 会覆盖源文件
	f0, err := os.Create("./goos/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	f0.Close()
	// openCloseFile()
	// createFile2()
	readOps()
}
