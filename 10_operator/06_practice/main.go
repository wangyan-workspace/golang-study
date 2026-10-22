package main

import "fmt"

func main() {
	//练习1：有两个变量，a和b，要求将其进行交换，最终打印结果
	a := 10
	b := 20
	temp := a
	a = b
	b = temp
	fmt.Printf("a=%v b=%v\n", a, b) //a=20 b=10

	//练习2：有两个变量，c和d，要求将其进行交换(不能使用中间变量)，最终打印结果
	c := 10
	d := 20
	c = c + d
	d = c - d
	c = c - d
	fmt.Printf("c=%v d=%v\n", c, d) //c=20 d=10

	// 练习 3：假如还有 100 天放假，问：xx 个星期零 xx 天
	var days int = 100
	var weeks = days / 7
	var day = days % 7
	fmt.Printf("距离放假还有 %v 个星期零 %v 天\n", weeks, day) // 距离放假还有 14 个星期零 2 天

	//练习 4：定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100), 请求出华氏温度对应的摄氏温度
	var F float64 = 100
	C := (F - 32) / 1.8
	fmt.Printf("华氏温度 %v 对应的摄氏温度为 %.2f\n", F, C) //华氏温度 100 对应的摄氏温度为 37.78

}
