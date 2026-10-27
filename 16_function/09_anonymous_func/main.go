package main

import "fmt"

func main() {
	//匿名函数  匿名自执行函数
	func() {
		fmt.Println("匿名函数") // 匿名函数
	}()

	//匿名函数
	var fn = func(x, y int) int {
		return x * y
	}
	fmt.Println(fn(10, 20)) // 200

	// 匿名自执行函数接收参数
	func(x, y int) {
		fmt.Println(x + y) // 30
	}(10, 20)
}