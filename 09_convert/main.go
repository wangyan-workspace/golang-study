package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1、整型和整型之间的转换
	var a int8 = 100
	var b int16 = 40
	fmt.Println(int16(a) + b) // 140

	// 2、浮点型和浮点型之间的转换
	var c float32 = 20
	var d float64 = 40
	fmt.Println(float64(c) + d) // 60

	// 3、整型和浮点型之间的转换
	var e float32 = 20.23
	var f int = 40
	fmt.Println(e + float32(f)) // 60.23

	// 4、通fmt.Sprintf() 把其他类型转换成 String 类型
	// 注意：Sprintf 使用中需要注意转换的格式 int 为%d  float 为%f  bool 为%t   byte 为%c
	var num int = 20
	var floatNum float64 = 20.23
	var t bool = true
	var byteNum byte = 'a'

	str1 := fmt.Sprintf("%d", num)
	fmt.Printf("值：%v 类型：%T\n", str1, str1) // 值：20 类型：string

	str2 := fmt.Sprintf("%f", floatNum)
	fmt.Printf("值：%v 类型：%T\n", str2, str2) // 值：20.230000 类型：string

	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("值：%v 类型：%T\n", str3, str3) // 值：true 类型：string

	str4 := fmt.Sprintf("%c", byteNum)
	fmt.Printf("值：%v 类型：%T\n", str4, str4) // 值：a 类型：string

	// 5、通过strconv把其他类型转换成string类型
	/*
		FormatInt
		参数1：int64 的数值
		参数2：传值int类型的进制
	*/
	var i int = 20
	str5 := strconv.FormatInt(int64(i), 10)
	fmt.Printf("值：%v 类型：%T\n", str5, str5) // 值：20 类型：string

	/*
		参数 1：要转换的值
		参数 2：格式化类型 'f'（-ddd.dddd）、
			 'b'（-ddddp±ddd，指数为二进制）、
			 'e'（-d.dddde±dd，十进制指数）、
			 'E'（-d.ddddE±dd，十进制指数）、
			 'g'（指数很大时用'e'格式，否则'f'格式）、
			 'G'（指数很大时用'E'格式，否则'f'格式）。
		 参数 3: 保留的小数点 -1（不对小数点格式化）
		 参数 4：格式化的类型 传入 64  32
	*/
	var floatNum2 float32 = 20.231313
	str6 := strconv.FormatFloat(float64(floatNum2), 'f', 4, 32)
	fmt.Printf("值：%v 类型：%T\n", str6, str6) // 值：20.2313 类型：string

	str7 := strconv.FormatBool(true)       //没有任何意义
	fmt.Printf("值：%v 类型：%T\n", str7, str7) // 值：true 类型：string

	character := 'b' // 没有任何意义
	str8 := strconv.FormatUint(uint64(character), 10)
	fmt.Printf("值：%v 类型：%T\n", str8, str8) // 值：98 类型：string

	// 6、string类型转换成整型
	str9 := "123456"
	fmt.Printf("%v---%T\n", str9, str9) // 123456---string

	/*
		ParseInt
		参数1：string数据
		参数2：进制
		参数3：位数 32 64 16

	*/
	num1, _ := strconv.ParseInt(str9, 10, 64)
	fmt.Printf("%v---%T\n", num1, num1) // 123456---int64

	/*
		ParseFloat
		参数1：string数据
		参数2：位数 32 64
	*/

	str10 := "123456.333xxxx"
	num2, _ := strconv.ParseFloat(str10, 64)
	fmt.Printf("%v---%T\n", num2, num2) // 0---float64

	//不建议把string类型转换成bool型
	b1, _ := strconv.ParseBool("xxxxxxx") // string 转 bool
	fmt.Printf("值：%v 类型：%T\n", b1, b1)    // 值：false 类型：bool

	/*
		104(h)
		101(e)
		108(l)
		108(l)
		111(o)
		32( )
		24352(张)
		19977(三)
	*/
	s := "hello 张三"
	for _, r := range s { //rune
		fmt.Printf("%v(%c) \n", r, r)
	}
}
