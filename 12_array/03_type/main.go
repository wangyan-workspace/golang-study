package main

import "fmt"

func main() {
	/*
		值类型 ：改变变量副本值的时候，不会改变变量本身的值
		引用类型：改变变量副本值的时候，会改变变量本身的值
	*/
	//值类型:基本数据类型 和 数组都是值类型

	var a = 10
	b := a
	a = 20
	fmt.Println(a, b) // 20 10

	var arr = [...]int{1, 2, 3}
	arr2 := arr
	arr2[0] = 11
	fmt.Println(arr, arr2) // [1 2 3] [11 2 3]

	//引用类型: 切片
	var s1 = []int{4, 5, 6, 7}
	s2 := s1
	s2[0] = 22
	fmt.Println(s1, s2) // [22 5 6 7] [22 5 6 7]
}
