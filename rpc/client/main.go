package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	fmt.Println("get conn")

	// 停止 rpc server
	time.Sleep(time.Second * 10)

	req := ArithRequest{9, 2}
	var res ArithResponse

	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		// log.Fatalln("arith error: ", err)
		// error connection is shut down
		fmt.Println("arith error: ", err)
	}

	// 再启动 server conn 不会自愈
	time.Sleep(time.Second * 30)

	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		// error 退出
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
