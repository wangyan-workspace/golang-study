package main

import "fmt"

/*
	结构体嵌套
*/
type User struct {
	Username string
	Password string
	Address Address //表示User结构体嵌套Address结构体
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
	// main.User{Username:"zhangsan", Password:"123345", Address:main.Address{Name:"李四", Phone:"19866732114", City:"北京"}}
	fmt.Printf("%#v\n", user) 
}