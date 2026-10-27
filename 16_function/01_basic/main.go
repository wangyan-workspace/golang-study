package main

import "fmt"

//求两个数的和
func sum (a int, b int) int {
	sum := a + b
	return sum
}

//求两个数的差
func sub (a int, b int) int {
	sub := a - b
	return sub
}

//函数参数的简写
func shortParams(a, b int) int {
	sub := a - b
	return sub
}

// 函数的可变参数，可变参数是指函数的参数数量不固定。
// Go 语言中的可变参数通过在参数名后加...来标识
func sumFn(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sumFn2(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x

	for _, v := range y {
		sum +=v
	}

	return sum
}

func main() {
	sum1 := sum(12,3)
	fmt.Println(sum1) // 15

	a := 34
	b := 5
	sub := sub(a, b)
	fmt.Println(sub) // 29

	sub2 := shortParams(a, b)
	fmt.Println(sub2) // 29

	sum2 := sumFn(12,34,56,78,90)
	fmt.Println(sum2) // 270

	sum3 := sumFn(12,34)
	fmt.Println(sum3) // 46

	sum4 := sumFn2(45,67,2,23)
	fmt.Println(sum4) // 137
}