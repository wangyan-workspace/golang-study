package main

import "fmt"

func main() {
	//我们想在切片里面放一系列用户的信息,这时候我们就可以定义一个元素为map类型的切片
	var userInfo = make([]map[string]string, 3, 3)
	fmt.Println(userInfo) // [map[] map[] map[]]

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

	if(userInfo[2] == nil) {
		userInfo[2] = make(map[string]string)
		userInfo[2]["name"] = "wangwu"
		userInfo[2]["age"] = "22"
		userInfo[2]["sex"] = "male"
	}
	fmt.Println(userInfo) // [map[age:18 name:zhangsan sex:male] map[age:20 name:lisi sex:female] map[age:22 name:wangwu sex:male]]

	for _, v := range userInfo {
		for key, value := range v {
			/*
				key: age --- value: 18
				key: sex --- value: male
				key: name --- value: zhangsan
				key: name --- value: lisi
				key: age --- value: 20
				key: sex --- value: female
				key: name --- value: wangwu
				key: age --- value: 22
				key: sex --- value: male
			*/
			fmt.Printf("key: %v --- value: %v\n", key, value)
		}
	}
}