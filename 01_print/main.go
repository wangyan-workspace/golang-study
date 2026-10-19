package main

import "fmt"

func main() {
	// fmt.Println("go", "python", "php", "javascript") // go python php javascript
	// fmt.Print("go", "python", "php", "javascript")   // gopythonphpjavascript

	// aaabbbccc
	fmt.Print("aaa")
	fmt.Print("bbb")
	fmt.Print("ccc")

	fmt.Println("ddd") // ddd
	fmt.Println("eee") // eee
	fmt.Println("fff") // fff

	var a int = 10 //go语言中变量定义以后必须要使用
	var b int = 20
	var c int = 30
	fmt.Println("a=", a, "b=", b, "c=", c) // a= 10 b= 20 c= 30
	fmt.Printf("a=%d b=%d c=%d", a, b, c)  // a=10 b=20 c=30
}
