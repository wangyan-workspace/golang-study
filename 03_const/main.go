package main

import (
	"fmt"
	"math"
)

func main() {
	// go语言中的常量
	const pi = 3.1415926
	fmt.Println(pi) // 3.1415926

	const a = "aaa"
	// a = "bbb"      // 常量定义后不能修改
	fmt.Println(a) // aaa

	// 多个常量也可以一起声明

	const (
		A = "aaa"
		B = "bbb"
		C = "ccc"
	)
	fmt.Println(A, B, C) // aaa bbb ccc

	// const 同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3) // 100 100 100

	const mathpi = math.Pi
	fmt.Println(mathpi) // 3.141592653589793

	// iota 常量计数器，只能在常量的表达式中使用
	const num = iota
	fmt.Println(num) // 0

	// 每次 const 出现时，都会让 iota 初始化为 0.【自增长】
	const (
		v1 = iota
		v2
		v3
		v4
	)
	fmt.Println(v1, v2, v3, v4) // 0 1 2 3

	// const iota 使用_跳过某些值
	const (
		t1 = iota
		_  // 跳过
		t2
		t3
	)
	fmt.Println(t1, t2, t3) // 0 2 3

	// iota 声明中间插队
	const (
		x1 = iota
		x2 = 100
		x3 = iota
		x4
	)
	fmt.Println(x1, x2, x3, x4) // 0 100 2 3

	// 多个 iota 定义在一行
	const (
		k1, k2 = iota + 1, iota + 2 // 1, 2
		k3, k4                      // 2, 3
		k5, k6                      // 3, 4
	)
	fmt.Println(k1, k2, k3, k4, k5, k6) // 1 2 2 3 3 4
}
