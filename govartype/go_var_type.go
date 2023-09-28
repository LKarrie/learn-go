package govartype

import (
	"fmt"
	"math"
	"unsafe"
)

func GoVarType() string {

	// 布尔
	var b1 bool = true
	var b2 bool = false
	var b3 = true
	b4 := false

	// 逻辑判断中使用布尔
	int1 := 1
	ok := int1 > 0
	if ok {
		fmt.Printf("\"true\": %v\n", "true")
	}

	// 循环语句中使用
	count := 1
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}

	// 不能使用 数字 或者 nil 表示真假
	// if ( 1 ) { fmt.Printf("\"true\": %v\n", "true") }

	// 数字类型
	// 整数
	// 无符号整数（不包含负数）
	// 浮点数
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)

	fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
	fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)
	fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)
	fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

	var f32 float32
	var f64 float64

	fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64)

	var ui = uint(math.MaxUint64) //再+1会导致overflows错误
	fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui)

	var imax, imin int
	imax = int(math.MaxInt64) //再+1会导致overflows错误
	imin = int(math.MinInt64) //再-1会导致overflows错误

	fmt.Printf("%T %dB %v~%v\n", imax, unsafe.Sizeof(imax), imin, imax)

	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("b4: %v\n", b4)

	// 复数
	// var c1 complex64
	var c1 = 1 + 2i
	// var c2 complex128
	var c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	return "learn-go-var-type"
}
