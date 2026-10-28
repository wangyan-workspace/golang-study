package main

import "fmt"

//注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
type MyInt int

func (m MyInt) PrintInfo() {
	fmt.Printf("我是自定义类型里面的自定义方法")
}


func main() {
	var a MyInt = 10
	a.PrintInfo() // 我是自定义类型里面的自定义方法
}