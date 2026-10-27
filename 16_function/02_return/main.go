package main

import "fmt"

//return 关键词一次可以返回多个值
func calc(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

// 返回值命名: 函数定义时可以给返回值命名，
// 并在函数体中直接使用这些变量，
// 最后通过 return 关键字 返回。
func calc1(a int, b int) (sum int, sub int) {
	fmt.Println(sum, sub) // 0 0
	sum = a + b
	sub = a - b
	return
}

func calc2(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func test() {
	fmt.Println("执行...")
}

func main() {
	a, b := calc(10, 2)
	fmt.Println(a, b) // 12 8

	c, d := calc1(10, 2)
	fmt.Println(c, d) // 12 8

	e, _ := calc2(10, 2)
	_, f := calc2(10, 2)
	fmt.Println(e, f) // 12 8

	test()
}