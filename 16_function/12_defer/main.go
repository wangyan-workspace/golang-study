package main

import "fmt"

/*
	Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
	在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
	也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
*/

func fn1() {
	fmt.Println("start")
	defer func() {
		fmt.Println("aaa")
		fmt.Println("bbb")
	}()
	fmt.Println("end")
}

func fn2() int{
	var a int
	defer func() {
		a++
	}()
	fmt.Println("end")
	return a
}

func fn3() (a int) {
	defer func() {
		a++
	}()
	return a
}

func main() {
	//1、defer的使用演示
	/*
		start
		end
		defer3
		defer2
		defer1
	*/
	fmt.Println("start")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	fmt.Println("end")

	//2、defer在命名返回值和匿名返回函数中表现不一样
	/*
		start
		end
		aaa
		bbb
	*/
	fn1()

	fmt.Println(fn2()) // 0

	fmt.Println(fn3()) // 1
}