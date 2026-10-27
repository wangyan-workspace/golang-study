package main

import "fmt"

// 函数的递归调用: 传入1个整数，递归打印出1-到这个数之内的所有整数
func fn1(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}
}

// 递归实现1-100的和
func fn2(n int) int {
	if n > 1 {
		return n + fn2(n - 1)
	} else {
		return 1
	}
}

//递归实现5的阶乘
func fn3(n int) int {
	if n > 1 {
		return n * fn3(n - 1)
	} else {
		return 1
	}
}

func main() {
	//1、for循环实现1到100的和
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum) // 5050

	//2、传入1个整数，递归打印出1-到这个数之内的所有整数
	/*
		10
		9
		8
		7
		6
		5
		4
		3
		2
		1
	*/
	fn1(10)

	//3、递归实现1-100的和
	fmt.Println(fn2(100)) // 5050

	//4、递归实现5的阶乘
	fmt.Println(fn3(5)) // 120
}