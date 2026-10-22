package main

import "fmt"

func main() {
	/*
		==	检查两个值是否相等，如果相等返回 True 否则返回 False。
		!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。
		>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。
		>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。
		<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。
		<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。
	*/

	//演示关系运算符的使用
	var a7 = 9
	var a8 = 8
	fmt.Println(a7 == a8) //false
	fmt.Println(a7 != a8) //true
	fmt.Println(a7 > a8)  //true
	fmt.Println(a7 >= a8) //true
	fmt.Println(a7 < a8)  //flase
	fmt.Println(a7 <= a8) //flase
	flag := a7 > a8
	fmt.Println("flag=", flag) // flag= true

	var a9 = 9
	var a10 = 9

	if a9 == a10 {
		fmt.Println("a9 == a10")
	} else {
		fmt.Println("a9 != a10")
	}
}
