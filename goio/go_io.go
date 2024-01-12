package goio

import "fmt"

// scan
// \n 是空格
func test1() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scan(&age)
	fmt.Printf("age: %v\n", age)
}

// Scanf
// 可以读取 \n
func test2() {
	var name string
	fmt.Println("请输入姓名：")
	fmt.Scanf("%s", &name)
	fmt.Printf("name: %v\n", name)
}

// Scanln
// \n 结束
func test3() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Printf("age: %v\n", age)
}

func GoIo() {

	// Go 语言中 将 IO 操作封装在了如下几个包中
	// io 为 IO 原语（I/O primitives）提供基本的接口  os File Reader Writer
	// io/ioutil 封装一些实用的 I/O 函数
	// fmt 实现格式化 I/O，类似 C 语言中的 printf 和 scanf  format fmt
	// bufio 实现带缓冲I/O

	// 在 io 包中最重要的是两个接口：Reader 和 Writer 接口 只要实现了这两个接口它就有了 IO 的功能
	// 	type Reader interface {
	// 		Read(p []byte) (n int, err error)
	// 	}
	// 	type Writer interface {
	// 		Write(p []byte) (n int, err error)
	// 	}
	// os.File 同时实现了 io.Reader 和 io.Writer
	// strings.Reader 实现了 io.Reader
	// bufio.Reader/Writer 分别实现了 io.Reader 和 io.Writer
	// bytes.Buffer 同时实现了 io.Reader 和 io.Writer
	// bytes.Reader 实现了 io.Reader
	// compress/gzip.Reader/Writer 分别实现了 io.Reader 和 io.Writer
	// crypto/cipher.StreamReader/StreamWriter 分别实现了 io.Reader 和 io.Writer
	// crypto/tls.Conn 同时实现了 io.Reader 和 io.Writer
	// encoding/csv.Reader/Writer 分别实现了 io.Reader 和 io.Writer

	// ioutil 包
	// ReadAll 读取数据，返回读到的字节 slice
	// ReadDir 读取一个目录，返回目录入口数组 []os.FileInfo
	// ReadFile 读一个文件，返回文件内容（字节slice）
	// WriteFile 根据文件路径，写入字节slice
	// TempDir 在一个目录中创建指定前缀名的临时目录，返回新临时目录的路径
	// TempFile 在一个目录中创建指定前缀名的临时文件，返回 os.File

	// fmt包实现了格式化的I/O函数 这点类似C中的printf和scanf
	// test1()
	// test2()
	// test3()
}
