package main

import (
	"fmt"
	"sort"
)

func mapSort(m map[string]string) string {
	var sliceArr []string

	// 1、把map对象的key放在一个切片里面
	for k, _ := range m {
		sliceArr = append(sliceArr, k)
	}

	//2、对key进行升序排序
	sort.Strings(sliceArr)

	var str string
	for _, v := range sliceArr {
		str += fmt.Sprintf("%v=>%v ", v, m[v])
	}

	return str
}

func main() {
	/*
		案例2：
				m1 := map[string]string{
					"username":"zhangsan",
					"age":"20",
					"sex":"男",
					"height":"180",
				}
				输出：age=>20height=>180sex=>男username=>zhangsan

				上面有一个map对象m1,封装一个方法,要求按照key升序排序，最后输出一个key=>valuekey=>value的字符串
	*/

	m1 := map[string]string{
		"username":"zhangsan",
		"age":"20",
		"sex":"男",
		"height":"180",
	}

	str := mapSort(m1)
	fmt.Println(str) // age=>20 height=>180 sex=>男 username=>zhangsan
}