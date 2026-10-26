package main

import "fmt"

func main() {
	//如果我们想在map对象中存放一系列的属性的时候，我们就可以把map类型的值定义成切片
	var userInfo = make(map[string][]string)

	userInfo["hobby"] = []string{
		"reading",
		"writing",
		"coding",
	}

	userInfo["skill"] = []string{
		"go",
		"python",
		"java",
	}
	fmt.Println(userInfo) // map[hobby:[reading writing coding] skill:[go python java]]

	/*
		---------------
		reading
		writing
		coding
		---------------
		go
		python
		java
	*/
	for _, v := range userInfo {
		fmt.Println("---------------")
		for _, value := range v {
			fmt.Println(value)
		}
	}
}