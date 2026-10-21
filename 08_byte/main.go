package main

import "fmt"

func main() {
	// 1、golang中定义字符  字符属于int类型
	var a = 'a'
	fmt.Printf("值：%v  类型:%T\n", a, a) //值：97  类型:int32
	// 当我们直接输出 byte（字符）的时候输出的是这个字符对应的ASCII码值

	//2、原样输出字符
	var b = 'a'
	fmt.Printf("值：%c  类型:%T\n", b, b) //值：a  类型:int32

	//3、定义一个字符串输出字符串里面的字符
	var str = "this"
	fmt.Printf("值：%v 原样输出%c  类型:%T\n", str[2], str[2], str[2]) // 值：105 原样输出i  类型:uint8

	//4、一个汉子占用 3个字节(utf-8), 一个字母占用一个字节

	// unsafe.Sizeof() 没法查看string类型数据所占用的存储空间

	var str1 = "this"      // 占用4个字节
	fmt.Println(len(str1)) // 4

	var str2 = "你好go"
	fmt.Println(len(str2)) // 8

	// 5、定义一个字符 字符的值是汉子

	//golang中汉子使用的是utf8编码 编码后的值就是int类型

	var c = '国'
	fmt.Printf("值：%v  类型:%T\n", c, c) // 值：22269  类型:int32
	fmt.Printf("值：%c  类型:%T\n", c, c) // 值：国  类型:int32

	//6、通过循环输出字符串里面的字符
	s := "你好 golang"

	/*
		228(ä)
		189(½)
		160( )
		229(å)
		165(¥)
		189(½)
		32( )
		103(g)
		111(o)
		108(l)
		97(a)
		110(n)
		103(g)
	*/
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c)\n", s[i], s[i])
	}

	/*
		20320(你)
		22909(好)
		32( )
		103(g)
		111(o)
		108(l)
		97(a)
		110(n)
		103(g)
	*/
	for _, v := range s { //rune
		fmt.Printf("%v(%c)\n", v, v)
	}

	//7、修改字符串
	s1 := "big"
	buteStr := []byte(s1)
	buteStr[0] = 'p'
	fmt.Println(string(buteStr)) //pig

	s2 := "你好golang"
	runeStr := []rune(s2)
	runeStr[0] = '我'
	fmt.Println(string(runeStr)) //我好golang
}
