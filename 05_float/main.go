package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1、定义float类型
	var a float32 = 3.12
	fmt.Printf("值：%v---%f,类型%T\n", a, a, a) // 值：3.12---3.120000,类型float32
	fmt.Println(unsafe.Sizeof(a))           // 4 float32占用4个字节

	var b float64 = 3.12
	fmt.Printf("值：%v---%f,类型%T\n", b, b, b) // 值：3.12---3.120000,类型float64
	fmt.Println(unsafe.Sizeof(b))           // 8 float64占用8个字节

	//2、 %f  输出float类型   %.2f 输出数据的时候保留2位小数
	var c float64 = 3.1415925535
	fmt.Printf("%v---%f---%T\n", c, c, c) // 3.1415925535---3.141593---float64
	fmt.Printf("%.2f\n", c)               // 3.14

	// 保留4位小数
	var d float64 = 3.1415925535
	fmt.Printf("%.4f\n", d) // 3.1416

	// 3、64位的系统中 浮点数默认是 float64
	f1 := 3.1345656456
	fmt.Printf("%f---%T\n", f1, f1) // 3.134566---float64

	// 4、Golang 科学计数法表示浮点类型
	var f2 float32 = 3.14e2         //表示f2等于3.14*10的2次方
	fmt.Printf("%v---%T\n", f2, f2) // 314---float32

	var f3 float64 = 3.14e-2        //表示f3等于3.14*10的-2次方(即3.14除以10的2次方)
	fmt.Printf("%v---%T\n", f3, f3) // 0.0314---float64

	// 5、Golang 中 float 精度丢失问题
	var f4 float64 = 1129.6
	fmt.Println(f4 * 100) // 112959.99999999999

	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2) //期望是 4.4，结果打印出了 4.3999999999999995

	// 6、int类型转换成float类型
	a1 := 10
	b1 := float64(a1)
	fmt.Printf("a1的类型是%T,b1的类型是%T\n", a1, b1) // a1的类型是int,b1的类型是float64

	// 7、float类型转换成float类型
	var a2 float32 = 23.4
	b2 := float64(a2)
	fmt.Printf("a2的类型是%T,b2的类型是%T\n", a2, b2) // a2的类型是float32,b2的类型是float64

	// 8、float类型转换成int类型(不建议这样做)
	var a3 float32 = 23.45
	b3 := int(a3)
	fmt.Printf("a3的类型是%T,b3的类型是%T\n", a3, b3) // a3的类型是float32,b3的类型是int
}
