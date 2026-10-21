package main

import (
	"fmt"
	"strings"
)

func main() {
	//1、定义string类型
	var str1 string = "hello"
	var str2 = "world"
	str3 := "你好"
	fmt.Printf("%v---%T\n", str1, str1) // hello---string
	fmt.Printf("%v---%T\n", str2, str2) // world---string
	fmt.Printf("%v---%T\n", str3, str3) // 你好---string

	//2、字符串转义符
	str4 := "hello\tworld\n"
	fmt.Println(str4) // hello	world

	str5 := "C:\\Go\\bin" // \t制表符，\n换行符，\ 反斜杠
	fmt.Println(str5)     // C:\Go\bin

	/*
		C:Go"bin
	*/
	str6 := "C:Go\"bin"
	fmt.Println(str6) // C:Go"bin

	//3、多行字符串
	str7 := `this is str1
	this is str2
	this is str3
	this is str4
	this is str5
	`
	fmt.Println(str7)

	// 4、len(str) 求长度
	var s1 = "hello"     // 长度5
	fmt.Println(len(s1)) // 5

	var s2 = "你好"        // 长度6
	fmt.Println(len(s2)) // 6

	// 5、+ 或者 fmt.Sprintf拼接字符串
	str11 := "hello"
	str12 := "world"
	str13 := str11 + str12
	fmt.Println(str13) // helloworld

	str14 := "hello"
	str15 := "world"
	str16 := fmt.Sprintf("%v %v", str14, str15)
	fmt.Println(str16) // hello world

	str17 := "hello" + "world" + "你好"
	fmt.Println(str17) // helloworld你好

	//6、strings.Split 分割字符串  strings需要引入strings包
	str18 := "123-456-789"
	arr := strings.Split(str18, "-")
	fmt.Println(arr) // [123 456 789]

	//7、strings.Join(a[]string, sep string) join 操作  表示把切片链接成字符串
	str19 := "123-456-789"
	arr1 := strings.Split(str19, "-")
	str20 := strings.Join(arr1, "*")
	fmt.Println(str20) // 123*456*789

	arr2 := []string{"123", "456", "789"}
	fmt.Println(arr2) // [123 456 789]
	str := strings.Join(arr2, "&")
	fmt.Println(str) // 123&456&789

	// 8、 strings.contains 判断是否包含
	ss := "this is str"
	ss1 := "is"
	flag := strings.Contains(ss, ss1)
	fmt.Println(flag) // true

	//9、strings.HasPrefix,strings.HasSuffix 前缀/后缀判断
	ss2 := "this is str"
	ss3 := "this"
	falg1 := strings.HasPrefix(ss2, ss3)
	fmt.Println(falg1) // true

	ss4 := "str"
	flag2 := strings.HasSuffix(ss2, ss4)
	fmt.Println(flag2) // true

	// 10、 strings.Index(),strings.LastIndex() 子串出现的位置  查找不到返回-1 查找到返回下标位置 下标是从0开始的
	ss5 := "this is str"
	ss6 := "is"
	num := strings.Index(ss5, ss6)
	fmt.Println(num) // 2

	ss7 := "s"
	num1 := strings.LastIndex(ss5, ss7)
	fmt.Println(num1) // 8

	ss8 := "is"
	num2 := strings.LastIndex(ss5, ss8)
	fmt.Println(num2) // 5

	ss9 := "xxx"
	num3 := strings.LastIndex(ss5, ss9)
	fmt.Println(num3) // -1 查找不到返回-1
}
