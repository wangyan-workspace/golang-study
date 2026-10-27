package main

import "fmt"

type calc func(int, int) int // 表示定义一个calc的类型

type myInt int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func test() {
	fmt.Println("test...")
}

func main() {
	var c calc
	c = sub
	fmt.Printf("c的类型：%T\n", c) // c的类型：main.calc
	fmt.Println(c(10, 5)) // 5

	d := sub
	fmt.Printf("d的类型：%T\n", d) // d的类型：func(int, int) int
	fmt.Println(d(10, 5)) // 5

	var a int = 10
	var b myInt = 20
	fmt.Printf("a的类型：%T\n", a) // a的类型：int
	fmt.Printf("b的类型：%T\n", b) // b的类型：main.myInt
	// fmt.Println(a + b) // 这里目前会报错，因为类型不一致，需要转换相同类型才能相加
	fmt.Println(a + int(b)) // 30 需要转换相同类型才能相加
}