package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

func main() {
	// Golang中空接口和类型断言使用细节
	var userInfo = make(map[string]interface{})
	userInfo["username"] = "张三"
	userInfo["age"] = 20
	userInfo["hobby"] = []string{"篮球", "足球", "乒乓球"}

	fmt.Println(userInfo["age"])   //20
	fmt.Println(userInfo["hobby"]) //[篮球 足球 乒乓球]
	// fmt.Println(userInfo["hobby"][1]) // interface {} does not support indexing

	var address = Address{
		Name:  "李四",
		Phone: 15246897564,
	}
	fmt.Println(address.Name) //李四

	userInfo["address"] = address
	fmt.Println(userInfo["address"]) //{李四 15246897564}

	// var name = userInfo["address"].Name // //type interface {} is interface with no methods
	// fmt.Println(name) //李四

	hobby2, _ := userInfo["hobby"].([]string)
	fmt.Println(hobby2[1]) //足球

	address2, _ := userInfo["address"].(Address)
	fmt.Println(address2.Name) //李四
}
