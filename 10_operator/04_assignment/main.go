package main

import "fmt"

func main() {
	/*
		=	简单的赋值运算符，将一个表达式的值赋给一个左值
		+=	相加后再赋值
		-=	相减后再赋值
		*=	相乘后再赋值
		/=	相除后再赋值
		%=	求余后再赋值
	*/
	var a = 20
	a = 21
	fmt.Println(a) // 21

	var a1 = 23 + 2
	fmt.Println(a1) // 25

	var a2 = 10
	b := a2
	fmt.Println(b) // 10

	var a3 = 10
	b1 := a3 + 2
	fmt.Println(b1) // 12

	var a4 = 10
	a4 += 3         // 等价于a4=a4+3
	fmt.Println(a4) // 13

	var a5 = 10
	a5 = a5 + 3
	fmt.Println(a5) // 13

	var a6 = 10
	a6 -= 3         //a6=a6-3
	fmt.Println(a6) // 7

	var a7 = 10
	a7 *= 3         //a7=a7*3
	fmt.Println(a7) // 30

	var a8 = 10
	a8 /= 3         // a8=a8/3
	fmt.Println(a8) //结果：3

	var a9 float64 = 10
	a9 /= 3         // a=a9/3
	fmt.Println(a9) //结果：3.3333333333333335

	var a10 = 10
	a10 %= 3         // a=a%3
	fmt.Println(a10) //1
}
