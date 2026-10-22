package main

import "fmt"

func main() {
	/*
		+	相加
		-	相减
		*	相乘
		/	相除
		%	求余=被除数-（被除数/除数）*除数

		1、除法注意：如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
		2、取余注意  余数=被除数-（被除数/除数）*除数
		3、注意： ++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。
	*/
	var a = 6
	var b = 3
	fmt.Println(a + b) // 9
	fmt.Println(a - b) // 3
	fmt.Println(a * b) // 18
	var c = a * b
	fmt.Println(c) // 18

	// 1、除法注意：如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	var a1 = 10
	var b1 = 3
	fmt.Println(a1 / b1) // 3

	var a2 = 10.0
	var b2 = 3.0
	fmt.Println(a2 / b2) // 3.3333333333333335

	// 2、取余注意    余数=被除数-（被除数/除数）*除数
	var a3 = 10
	var b3 = 3
	fmt.Println(a3 % b3) //1

	fmt.Println(-10 % 3) // - 10 - (-10/3)*3  =-1

	fmt.Println(10 % -3) // 10-(10/-3)*-3=1

	//3、注意： ++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。
	//注意：在 golang 中，++ 和 -- 只能独立使用 错误写法如下：
	/*
		var i int = 8
		var a4 int
		a4 = i++ //错误，i++只能独立使用
		a4 = i-- //错误, i--只能独立使用
		fmt.Ptintln(a4)
	*/

	//注意：在 golang 中没有前++ 错误写法如下：
	// var a5 = 12
	// ++a
	// fmt.Println(a5)

	// 正确的写法
	var a6 = 12
	a6++
	a6--
	fmt.Println(a6) // 12
}
