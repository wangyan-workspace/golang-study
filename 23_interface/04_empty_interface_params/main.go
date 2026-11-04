package main

import "fmt"

//golang中空接口也可以直接当做类型来使用，可以表示任意类型

// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("值:%v 类型:%T\n", a, a)
}

func main() {
	var a interface{}
	a = 20
	fmt.Printf("值:%v 类型:%T\n", a, a) // 值:20 类型:int
	a = "你好golang"
	fmt.Printf("值:%v 类型:%T\n", a, a) // 值:你好golang 类型:string
	a = true
	fmt.Printf("值:%v 类型:%T\n", a, a) // 值:true 类型:bool

	fmt.Println("--------------------")

	show(20)             // 值:20 类型:int
	show("hello golang") // 值:hello golang 类型:string
	slice := []int{1, 2, 3, 4}
	show(slice) // 值:[1 2 3 4] 类型:[]int

	fmt.Println("--------------------")

	var m1 = make(map[string]interface{})
	m1["username"] = "zhangsan"
	m1["age"] = 20
	m1["married"] = true
	fmt.Println(m1) // map[age:20 married:true username:zhangsan]

	fmt.Println("--------------------")

	var s1 = []interface{}{1, 2, "hello", true} // [1 2 hello true]
	fmt.Println(s1)
}
