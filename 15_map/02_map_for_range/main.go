package main

import "fmt"

func main() {
	userInfo := map[string]string{
		"name": "lisi",
		"age": "20",
		"sex": "male",
	}

	fmt.Println(userInfo) // map[age:20 name:lisi sex:male]

	for k, v := range userInfo {
		/*
			key:name value:lisi
			key:age value:20
			key:sex value:male
		*/
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}