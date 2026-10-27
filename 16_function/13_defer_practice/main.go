package main

import "fmt"

func fn1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func fn2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func fn3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //5
}

func fn4() (x int) {
	defer func(x int) {
		// fmt.Println(x) // 0
		x++
	}(x) //defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	// fmt.Println(x) // 0
	return 5
}

func main() {
	fmt.Println(fn1()) // 5
	fmt.Println(fn2()) // 6
	fmt.Println(fn3()) // 5
	fmt.Println(fn4()) // 5
}