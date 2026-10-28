package main

import "fmt"

func main() {
	// var userInfo map[string]string
	// userInfo["username"] = "张三"
	// fmt.Println(userInfo) // 会报错的

	var userInfo = make(map[string]string)
	userInfo["username"] = "lisi"
	fmt.Println(userInfo) // map[username:lisi]


	// var a []int
	// a[0] = 1
	// fmt.Println(a) // 会报错的

	var a = make([]int, 4, 4)
	a[0] = 1
	fmt.Println(a) // [1 0 0 0]

	// var b *int //指针也是引用数据类型
	// *b = 100
	// fmt.Println(*b) // 报错
}