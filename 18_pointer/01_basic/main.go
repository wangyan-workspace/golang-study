package main

import "fmt"

func main() {
	//在计算机底层 a 这个变量其实对应了一个内存地址
	var a int = 10
	fmt.Printf("a的值是%v a的类型%T a的内存地址是%p\n", a, a, &a) // a的值是10 a的类型int a的内存地址是0x14000104020

	//每一个变量都有自己的内存地址
	// 指针也是一个变量，但它是一种特殊的变量，它存储的数据不是一个普通的值，而是另 一个变量的内存地址
	var b = 20
	var p = &b // p指针变量   p的类型 *int（指针类型）
	fmt.Printf("b的值是%v b的类型%T b的内存地址是%p\n", b, b, &b) // b的值是20 b的类型int b的内存地址是0x14000112028
	fmt.Printf("p的值是%v p的类型%T p的内存地址是%p\n", p, p, &p) // p的值是0x14000112028 p的类型*int p的内存地址是0x1400011c038
}