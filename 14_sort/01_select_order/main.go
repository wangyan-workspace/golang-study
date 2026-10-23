package main

import "fmt"

func main() {
	/*
		选择排序：进行从小到大排序
		概念:通过比较 首先选出最小的数放在第一个位置上;
			然后在其余的数中选出次小数放在第二个位置上;
		    依此类推,直到所有的数成为有序序列。
	*/
	/*

		第一轮:
			9 8 7 6 5 4
			8 9 7 6 5 4
			7 9 8 6 5 4
			6 9 8 7 5 4
			5 9 8 7 6 4
			4 9 8 7 6 5    4放在了正确的位置
		第二轮:
			  9 8 7 6 5
			  8 9 7 6 5
			  7 9 8 6 5
			  6 9 8 7 5
			  5 9 8 7 6    5放在了正确的位置

		第三轮:
				9 8 7 6
				8 9 7 6
				7 9 8 6
				6 9 8 7    6放在了正确的位置

		第四轮:
				  9 8 7
				  8 9 7
				  7 9 8    7放在了正确的位置
		第五轮:
					9 8
					8 9    8放在了正确的位置
	*/

	// 从小到大的排序
	var arr = []int{6, 5, 4, 9, 8, 7}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr) // [4 5 6 7 8 9]

	// 从大到小的排序
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr) // [9 8 7 6 5 4]
}
