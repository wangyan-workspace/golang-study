package main

import "fmt"

/*
	嵌套匿名结构体
*/
type User struct {
	Username string
	Password string
	Address 
}

type Address struct {
	Name string
	Phone string
	City string
}
func main() {
	var user User
	user.Username = "zhangsan"
	user.Password = "123345"
	user.Address.Name = "李四"
	user.Address.Phone = "19866732114"
	user.Address.City = "北京"

	user.City = "上海" //当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。

	// main.User{Username:"zhangsan", Password:"123345", Address:main.Address{Name:"李四", Phone:"19866732114", City:"上海"}}
	fmt.Printf("%#v\n", user) 

	fmt.Println(user.Address.Phone) // 19866732114
	fmt.Println(user.Phone) // 19866732114
}