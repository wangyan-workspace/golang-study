package main

import "fmt"

func main() {
	//我们想在切片里面放一系列用户的信息,这时候我们就可以定义一个元素为map类型的切片
	var userInfo = make([]map[string]string, 2, 2)
	fmt.Println(userInfo) // [map[] map[]]
  fmt.Println(userInfo[0]) // map[] map不初始化的默认值nil
	fmt.Println(userInfo[1] == nil) // true


	if(userInfo[0] == nil) {
		userInfo[0] = make(map[string]string)
		userInfo[0]["name"] = "zhangsan"
		userInfo[0]["age"] = "18"
		userInfo[0]["sex"] = "male"
	}

	if(userInfo[1] == nil) {
		userInfo[1] = make(map[string]string)
		userInfo[1]["name"] = "lisi"
		userInfo[1]["age"] = "20"
		userInfo[1]["sex"] = "female"
	}
	fmt.Println(userInfo) // [map[age:18 name:zhangsan sex:male] map[age:20 name:lisi sex:female]]
}