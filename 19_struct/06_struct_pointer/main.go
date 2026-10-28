package main

import "fmt"

type Person struct {
	Name string
	Age int
	Sex string
	Height int
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%v 年龄:%v\n", p.Name, p.Age)
}

func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func main() {
	var p1  = Person{
		Name: "zhangsan",
		Age: 20,
		Sex: "男",
	}
	p1.PrintInfo() // 姓名：zhangsan 年龄:20

	p1.SetInfo("lisi",34)
	p1.PrintInfo() // 姓名：lisi 年龄:34
}