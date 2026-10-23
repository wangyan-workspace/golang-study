package main

import "fmt"

func main() {
	var arr = [3]int{1, 2, 3} // 一维数组
	fmt.Println(arr)          // [1 2 3]

	//1、二维数组的定义
	var arr1 = [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(arr1)       // [[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(arr1[0])    // [北京 上海]
	fmt.Println(arr1[2][1]) // [2][1]表示二维数组下标为2的元素中的下标为1的元素，即重庆

	/*
		北京
		上海
		广州
		深圳
		成都
		重庆
	*/
	for _, val := range arr1 {
		for _, val2 := range val {
			fmt.Println(val2)
		}
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Println(arr1[i][j])
		}
	}

	//定义二维数组的另一种方式
	var arr2 = [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(arr2) // [[北京 上海] [广州 深圳] [成都 重庆] [北京 上海] [广州 深圳] [成都 重庆]]
}
