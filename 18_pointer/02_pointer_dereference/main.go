package main

import "fmt"

func main() {
	var a = 10
	p := &a  //p指针变量  类型 *int

	// *p ：表示取出p这个变量对应的内存地址的值
	fmt.Println(p) // 0x14000102020   a的地址
	fmt.Println(*p) // 10  *p 表示打印a对应的值 

	*p = 30 // 改变p这个变量对应的内存地址的值
	fmt.Println(a) // 30
}