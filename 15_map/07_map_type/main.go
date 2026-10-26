package main

import "fmt"

func main() {
	/*
		值类型 ：改变变量副本值的时候，不会改变变量本身的值  (基本数据类型、数组)
		引用类型：改变变量副本值的时候，会改变变量本身的值 （切片、map）
	*/

	// map类型也是引用数据类型
	var userInfo = make(map[string]string)
	userInfo["name"] = "zhangsan"
	userInfo["age"] = "18"

	userInfo2 := userInfo
	userInfo2["name"] = "lisi"
	fmt.Println(userInfo) // map[age:18 name:lisi]
	fmt.Println(userInfo2) // map[age:18 name:lisi]
}