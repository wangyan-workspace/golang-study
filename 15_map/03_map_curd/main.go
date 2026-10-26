package main

import "fmt"

func main() {
	//map类型数据的curd
	//1、创建map类型的数据
	var userInfo = make(map[string]string)
	userInfo["name"] = "lisi"
	userInfo["age"] = "20"
	fmt.Println(userInfo) // map[age:20 name:lisi]

	//2、修改map类型的数据
	var userInfo2 = map[string]string{
		"name": "zhangsan",
		"age": "18",
		"sex": "female",
	}
	userInfo2["name"] = "wangwu"
	fmt.Println(userInfo2) // map[age:18 name:wangwu sex:female]

	//3、获取/查找map类型的数据
	var userInfo3 = map[string]string{
		"name": "lucky",
		"age": "18",
		"sex": "famale",
	}
	fmt.Println(userInfo3) //map[age:18 name:lucky sex:famale]

	value, ok := userInfo3["name"]
	fmt.Println(value, ok) // lucky true

	value2, ok2 := userInfo3["xxx"]
	fmt.Println(value2, ok2) // 空 false

	//4、删除map数据里面的key以及对于的值
	var userInfo4 = map[string]string{
		"name": "John",
		"age": "20",
		"sex": "male",
	}
	fmt.Println(userInfo4) // map[age:20 name:John sex:male]

	delete(userInfo4, "sex")
	fmt.Println(userInfo4) // map[age:20 name:John]
}