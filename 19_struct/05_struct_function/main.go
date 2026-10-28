package main

import "fmt"

/*
	结构体可以描述现实生活中的任何事物，生活中的任何事物都可以当做结构体对象
	我们可以把客观事物封装成结构体:
		汽车
			汽车有属性：颜色   大小   重量  发动机   轮胎 ...
			汽车行为 也叫方法：  run  跑

		小狗
			属性：颜色     大小     品种  性别 ..
			行为：叫    、闻一闻  舔一舔  咬一咬

		电风扇
			属性： 颜色 大小  高低...
			行为：转动

		人也是一个结构体对象：
			属性： 名字  年龄 性别 ...
			行为：工作  运动 ..
*/
type Person struct {
	Name string
	Age int
	Sex string
	Height int
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%v 年龄:%v\n", p.Name, p.Age)
}

func main() {
	var p1  = Person{
		Name: "zhangsan",
		Age: 20,
		Sex: "男",
	}
	p1.PrintInfo() // 姓名：zhangsan 年龄:20

	var p2 = Person{
		Name: "lisi",
		Age: 34,
		Sex: "女",
	}
	p2.PrintInfo() // 姓名：lisi 年龄:34

	p1.PrintInfo() // 姓名：zhangsan 年龄:20
}