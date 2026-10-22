package main

import "fmt"

func main() {
	// goto 语句通过标签进行代码间的无条件跳转。
	// goto 语句可以在快速跳出循环、避免重复退出上有一定的帮助。
	// Go 语言中使用 goto 语句能简化一些代码的实现过程

	/*
		成年人
		ccc
		ddd
	*/
	var n = 30
	if n > 24 {
		fmt.Println("成年人")
		goto label3
	}

	fmt.Println("aaa")
	fmt.Println("bbb")

label3:
	fmt.Println("ccc")
	fmt.Println("ddd")
}
