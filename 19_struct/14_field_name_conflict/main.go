package main

import "fmt"

/*
	嵌套匿名结构体
*/
type User struct {
	Username string
	Password string
	Address 
	Email
}

type Address struct {
	Name string
	Phone string
	City string
	AddTime string
}

type Email struct {
	Account string
	AddTime string
}

func main() {
	var user User
	user.Username = "zhangsan"
	user.Password = "123345"
	user.Address.Name = "李四"
	user.Address.Phone = "19866732114"
	user.Address.City = "北京"
	user.City = "上海"

	user.Address.AddTime = "2026-09-28"

	user.Email.AddTime = "2026-08-28"
  // main.User{Username:"zhangsan", Password:"123345", Address:main.Address{Name:"李四", Phone:"19866732114", City:"上海", AddTime:"2026-09-28"}, Email:main.Email{Account:"", AddTime:"2026-08-28"}}
	fmt.Printf("%#v\n", user) 
}