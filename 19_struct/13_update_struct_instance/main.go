package main

import "fmt"

/*
	嵌套匿名结构体
*/
type User struct {
	Username string
	Password string
	AddTime string
	Address 
}

type Address struct {
	Name string
	Phone string
	City string
	AddTime string
}
func main() {
	var user User
	user.Username = "zhangsan"
	user.Password = "123345"
	user.Address.Name = "李四"
	user.Address.Phone = "19866732114"
	user.Address.City = "北京"

	user.City = "上海" //当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。

	user.AddTime = "2026-10-28"

	user.Address.AddTime = "2026-09-28"

	// main.User{Username:"zhangsan", Password:"123345", AddTime:"2026-10-28", Address:main.Address{Name:"李四", Phone:"19866732114", City:"上海", AddTime:"2026-09-28"}}
	fmt.Printf("%#v\n", user) 
}