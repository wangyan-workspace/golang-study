package main

import "fmt"

func test() bool {
	fmt.Println("test...")
	return true
}

func main() {
	/*
		&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。
		||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。
		!	逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。
	*/
	var n = 23
	var v = 8
	fmt.Println(n > 10 && v < 10) // true
	fmt.Println(n > 24 && v < 10) // false
	fmt.Println(n > 5 && v < 6)   // false
	fmt.Println(n == 5 && v < 6)  // false

	var n1 = 23
	var v1 = 8
	fmt.Println(n1 > 10 || v1 < 10) // true
	fmt.Println(n1 > 24 || v1 < 10) // true
	fmt.Println(n1 > 5 || v1 < 6)   // true
	fmt.Println(n1 == 5 || v1 < 6)  // false

	flag1 := false
	fmt.Println(!flag1) // true

	//逻辑与和逻辑或短路
	//逻辑与：前面是false后面就不会执行  逻辑或：前面是true后面就不会执行

	/*
		test...
		执行
	*/
	var n2 = 10
	if n2 > 9 && test() {
		fmt.Println("执行")
	}

	/*
		输出：
		空
	*/
	var n3 = 10
	if n3 > 11 && test() {
		fmt.Println("执行")
	}

	/*
		输出：
		test...
		执行
	*/
	var n4 = 10
	if n4 > 11 || test() {
		fmt.Println("执行")
	}

	/*
		输出：执行
	*/
	var n5 = 10
	if n5 < 11 || test() {
		fmt.Println("执行")
	}
}
