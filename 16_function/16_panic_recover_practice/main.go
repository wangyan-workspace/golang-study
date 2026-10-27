package main

import "fmt"

func fn1(a int, b int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	fmt.Println(a / b)
}

func main() {
	/*
		err: runtime error: integer divide by zero
		main end
		5
	*/
	fn1(10, 0)
	fmt.Println("main end")
	fn1(10, 2)
}