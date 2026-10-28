package main

import "fmt"

// 注意：结构体首字母可以大写也可以小写，大写表示这个结构体是公有的，
// 在其他的包里面 可以使用。小写表示这个结构体是私有的，只有这个包里面才能使用
type Person struct {
	Name string
	Age int
	Sex string
}

func main() {
	// 注意：在 Golang 中支持对结构体指针直接使用.来访问结构体的成员。
	// p2.name = "张三" 其 实在底层是(*p2).name = "张三"
	var p2 = new(Person)
	p2.Name = "lisi"
	p2.Age = 20
	p2.Sex = "男"
	// 值:&main.Person{Name:"lisi", Age:20, Sex:"男"} 类型:*main.Person
	fmt.Printf("值:%#v 类型:%T\n", p2, p2)  

	(*p2).Name = "wangwu"
	// 值:&main.Person{Name:"wangwu", Age:20, Sex:"男"} 类型:*main.Person
	fmt.Printf("值:%#v 类型:%T\n", p2, p2)  


	var p3 = &Person{}
	p3.Name = "zhaosi"
	p3.Age = 56
	p3.Sex = "男"
	// 值:&main.Person{Name:"zhaosi", Age:56, Sex:"男"} 类型:*main.Person
	fmt.Printf("值:%#v 类型:%T\n", p3, p3)  


	var p4 = Person{
		Name: "zhangsan",
		Age: 18,
		Sex: "男",
	}
	// 值:main.Person{Name:"zhangsan", Age:18, Sex:"男"} 类型:main.Person
	fmt.Printf("值:%#v 类型:%T\n", p4, p4)  


	var p5 = &Person{
		Name: "zhaoliu",
		Age: 40,
		Sex: "女",
	}
	// 值:&main.Person{Name:"zhaoliu", Age:40, Sex:"女"} 类型:*main.Person
	fmt.Printf("值:%#v 类型:%T\n", p5, p5)  


	var p6 = &Person{
		Name: "wanghaha",
	}
	// 值:&main.Person{Name:"wanghaha", Age:0, Sex:""} 类型:*main.Person
	fmt.Printf("值:%#v 类型:%T\n", p6, p6)  

	var p7 = &Person{
		"zhangsan",
		23,
		"男",
	}
	// 值:&main.Person{Name:"zhangsan", Age:23, Sex:"男"} 类型:*main.Person
	fmt.Printf("值:%#v 类型:%T\n", p7, p7)  
}