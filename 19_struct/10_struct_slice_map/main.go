package main

import "fmt"

/*
	结构体的字段类型可以是：基本数据类型、也可以是切片、Map 以及结构体
	如果结构体的字段类型是: 指针，slice，和map的零值都是 nil ，即还没有分配空间
	如果需要使用这样的字段，需要先make，才能使用.
*/
type Person struct {
	Name string
	Age int
	Hobby []string
	Info map[string]string
}

func main() {
	var p Person
	p.Name = "zhangsan"
	p.Age = 20
	p.Hobby = make([]string, 3, 6)
	// p.Hobby = []string{"篮球", "足球", "乒乓球"} // make后可以直接赋值，与下方方式一样的
	p.Hobby[0] = "篮球"
	p.Hobby[1] = "足球"
	p.Hobby[2] = "乒乓球"
	p.Info = make(map[string]string)
	p.Info["address"] = "北京"
	p.Info["phone"] = "1234567890"

	// main.Person{Name:"zhangsan", Age:20, Hobby:[]string{"篮球", "足球", "乒乓球"}, Info:map[string]string{"address":"北京", "phone":"1234567890"}}
	fmt.Printf("%#v\n", p)
	// [篮球 足球 乒乓球]
	fmt.Printf("%v", p.Hobby)
}