package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1、定义int类型
	var num int = 10
	num = 20
	fmt.Printf("num=%v类型:%T", num, num) // num=20类型:int

	//2、int8的范围演示
	var num1 int8 = 98
	fmt.Printf("num1=%v类型:%T", num1, num1) // num1=98类型:int8

	// (-128 到 127)  错误
	// var num2 int8 = 130 // //constant 130 overflows int8
	// fmt.Printf("num2=%v类型:%T", num2, num2) // num2=130类型:int8

	var num3 int16 = 130
	fmt.Printf("num3=%v类型:%T", num3, num3) // num3=130类型:int16

	// 3、uint8的范围(0-255)
	// var n1 uint8 = -2 // //constant -2 overflows uint8
	// fmt.Printf("n1=%v类型:%T", n1, n1) // n1=-2类型:int

	var n2 uint8 = 200
	fmt.Printf("n2=%v类型:%T", n2, n2) // n2=200类型:uint8

	//4、int8 int16 ... 占用的存储空间大小  unsafe.Sizeof 查看不同长度的整型 在内存里面的存储空间
	var a int8 = 15
	fmt.Printf("a=%v类型:%T 占用的存储空间大小:%d", a, a, unsafe.Sizeof(a)) // a=15类型:int8 占用的存储空间大小:1

	var b int32 = 15
	fmt.Printf("b=%v类型:%T 占用的存储空间大小:%d", b, b, unsafe.Sizeof(b)) // b=15类型:int32 占用的存储空间大小:4

	// 5、int类型转换
	var a1 int32 = 10
	var a2 int64 = 21
	fmt.Println(int64(a1) + a2) //把a1转换成64位  31
	fmt.Println(a1 + int32(a2)) //把a2转换成32位  31

	//高位向底位转换的时候要注意
	var v1 int16 = 130
	fmt.Println(int8(v1)) //-126  有问题了

	var v2 int16 = 110
	fmt.Println(int8(v2))  // 110
	fmt.Println(int64(v2)) // 110

	//6、数字字面量语法  %d 表示10进制输出 %b表示二进制输出 %o 八进制输出 %x 表示16进制输出(了解)
	c := 12

	fmt.Printf("c=%v\n", c) //%v 原样输出 c=12

	fmt.Printf("c=%d\n", c) //%d 表示10进制输出 c=12

	fmt.Printf("c=%b\n", c) //%b 表示二进制输出 c=1100

	fmt.Printf("c=%o\n", c) //%o 表示八进制输出 c=14

	fmt.Printf("c=%x\n", c) //%x 表示16进制输出 c=c
}
