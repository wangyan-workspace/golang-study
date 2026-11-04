package main

import "fmt"

// 定义一个方法，可以传入任意数据类型，然后根据不同的类型实现不同的功能
func MyPrint1(x interface{}) {
	if _, ok := x.(string); ok {
		fmt.Println("string类型")
	} else if _, ok := x.(int); ok {
		fmt.Println("int类型")
	} else if _, ok := x.(bool); ok {
		fmt.Println("bool类型")
	}
}

func MyPrint2(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("传入错误...")
	}
}

func main() {
	var a interface{}
	a = "hello golang"
	v, ok := a.(string)
	if ok {
		fmt.Println("a就是一个string类型,值是：", v)
	} else {
		fmt.Println("断言失败")
	}

	MyPrint1("hello golang") // string类型
	MyPrint1(true)           // bool类型

	MyPrint2("hello golang") // string类型
	MyPrint2(true)           // bool类型
}
