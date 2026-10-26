package main

import "fmt"

func main() {
	// 1、make创建map类型的数据
	var userInfo = make(map[string]string)
	userInfo["name"] = "lisi"
	userInfo["age"] = "20"
	userInfo["sex"] = "male"
	fmt.Println(userInfo)         // map[age:20 name:lisi sex:male]
	fmt.Println(userInfo["name"]) // lisi

	//2、map 也支持在声明的时候填充元素
	var userInfo2 = map[string]string{
		"name": "zhangsan",
		"age": "18",
		"sex": "female",
	}
	fmt.Println(userInfo2) // map[age:18 name:zhangsan sex:female]

	// 3、第三种创建map类型数据的方法
	userInfo3 := map[string]string{
		"name": "wangwu",
		"age": "67",
		"sex": "male",
	}
	fmt.Println(userInfo3) // map[age:67 name:wangwu sex:male]
}
