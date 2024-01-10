package concurrency

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func show(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v,%v\n", msg, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func reponseSize(url string) {
	println("Step1: ", url)
	reponse, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	println("Step2: ", url)
	defer reponse.Body.Close()

	println("Step3: ", url)
	body, err := ioutil.ReadAll(reponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	println("Step4: ", len(body))

}

func GoGoroutines() {
	// Go 中的并发是函数相互独立运行的能力。
	// Goroutines是并发运行的函数 Go提供了Goroutines作为并发处理操作的一种方式

	// 创建协程非常简单，就是在一个任务函数前添加go关键字
	// go task()

	go show("golang")
	// 在main协程中执行，如果它前面添加go 如果不在底部Sleep 程序无输出 因为main 已经提前结束了
	go show("golang")
	println("end ...")

	go reponseSize("https://baidu.com")
	go reponseSize("https://jd.com")

	time.Sleep(3 * time.Second)
}
