package main

import "fmt"

func main() {
	//1、数组的长度是类型的一部分
	var arr1 [5]int
	var arr2 [4]int
	var strArr [5]string
	fmt.Printf("arr1: %T arr2: %T strArr: %T\n", arr1, arr2, strArr) // arr1: [5]int arr2: [4]int strArr: [5]string

	//2、数组的初始化 第一种方法
	var arr3 [3]int
	arr3[0] = 34
	arr3[1] = 23
	arr3[2] = 89
	fmt.Println(arr3) // [34 23 89]

	var strArr2 [3]string
	strArr2[0] = "php"
	strArr2[1] = "nodejs"
	strArr2[2] = "golang"
	// strArr2[3] = "java" // //out of bounds for 3-element array
	fmt.Println(strArr2) // [php nodejs golang]

	//2、数组的初始化 第二种方法
	var arr4 = [3]string{"php", "nodejs", "golang"}
	fmt.Println(arr4) // [php nodejs golang]

	var arr5 = [3]int{4, 5, 7}
	fmt.Println(arr5) // [4 5 7]

	arr6 := [5]int{3, 4, 5, 6, 7}
	fmt.Println(arr6) // [3 4 5 6 7]

	//3、数组的初始化 第三种方法  一般情况下我们可以让编译器 根据初始值的个数自行推断数组的长度
	var arr7 = [...]int{34, 56, 78, 43, 57, 89}
	fmt.Println(arr7)      // [34 56 78 43 57 89]
	fmt.Println(len(arr7)) // 6

	//注意数组的长度
	arr8 := [...]string{"php", "nodejs", "golang", "java", "python"}
	// arr8[5] = "c++" // out of bounds for 5-element array
	fmt.Println(arr8)

	//改变数组里面的值
	arr9 := [...]string{"php", "nodejs", "golang", "java", "python"}
	arr9[3] = "c++"
	fmt.Println(arr9) // [php nodejs golang c++ python]

	// 4、数组的初始化 第四种方法   我们还可以使用指定索引值的方式来初始化数组
	arr10 := [...]int{1: 30, 4: 9, 8: 38}
	fmt.Println(len(arr10)) // 9
	fmt.Println(arr10)      // [0 30 0 0 9 0 0 0 38]

	//5、数组的循环遍历 for  for range

	/*
		1
		2
		3
		4
		5
		6
		7
	*/
	var arr11 = [...]int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(arr11); i++ {
		fmt.Println(arr11[i])
	}

	/*
		php
		nodejs
		golang
		java
		python
	*/
	strArr3 := [...]string{"php", "nodejs", "golang", "java", "python"}
	for i := 0; i < len(strArr3); i++ {
		fmt.Println(strArr3[i])
	}

	/*
		key:0 value:php
		key:1 value:nodejs
		key:2 value:golang
		key:3 value:java
		key:4 value:python
	*/
	strArr4 := [...]string{"php", "nodejs", "golang", "java", "python"}
	for key, val := range strArr4 {
		fmt.Printf("key:%v value:%v\n", key, val)
	}
}
