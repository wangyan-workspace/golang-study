package main

import "fmt"

func main() {
	/*
		Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

		注意：
		1.布尔类型变量的默认值为false
		2.Go 语言中不允许将整型强制转换为布尔型
		3.布尔型无法参与数值运算，也无法与其他类型进行转换
	*/

	var flag = true
	fmt.Printf("%v---%T\n", flag, flag) // true---bool

	//1.布尔类型变量的默认值为false
	var a bool
	fmt.Printf("%v\n", a) // false

	//2.string型变量的默认值为空
	var s string
	fmt.Printf("%v\n", s) // 空

	//3.int型变量的默认值为0
	var i int
	fmt.Printf("%v\n", i) // 0

	//4.float型变量的默认值为0
	var f float32
	fmt.Printf("%v\n", f) // 0

	//5、Go 语言中不允许将整型强制转换为布尔型
	// var a = 1
	// if a { // 错误写法
	// 	fmt.Println("true")
	// }

	//6.布尔型无法参与数值运算，也无法与其他类型进行转换
	// var s = "this is str"
	// if s { // 错误写法
	// 	fmt.Println("true")
	// }

	var f1 = false
	if f1 { // 正确写法
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
