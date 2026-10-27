package main

import "fmt"

var a = "全局变量"

func run() {
	var b = "局部变量"
	fmt.Println("run方法a=", a)
	fmt.Println("run方法b=", b)
	fmt.Println("run")
}

func main() {
	/*
		函数变量作用域：

			全局变量：全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效 （全局作用域）

			局部变量：局部变量是函数内部定义的变量， 函数内定义的变量无法在该函数外使用 （局部作用域）
	*/

	fmt.Println(a) // 全局变量
	// fmt.Println("b=", b) // 报错，b是局部变量，无法在main函数中使用

	run()

	//i就是局部变量 只能在for方法体内使用
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}
	// fmt.Println("i=", i) // 报错，i是局部变量，无法在main函数中使用

	//在本作用域相当于全局变量
	var flag = true
	if flag {
		fmt.Println("true")
	}

	//flag块作用域 局部变量
	if flag1 := true; flag1 {
		fmt.Println("true1")
	}
	// fmt.Println("flag1=", flag1) // 报错，flag1是局部变量，无法在main函数中使用
}