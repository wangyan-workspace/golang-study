package main

import "fmt"

func main() {
	// 1、请求出一个数组里面元素的和以及这些元素的平均值。分别用for和for-range实现
	var arr = [...]int{34, 56, 76, 2, 3}
	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("arr数组元素的和是:%v\n", sum)                       // arr数组元素的和是:171
	fmt.Printf("平均值：%.2f\n", float64(sum)/float64(len(arr))) // 平均值：34.20

	var sum1 = 0
	for _, v := range arr {
		sum1 += v
	}
	fmt.Printf("arr数组元素的和是:%v\n", sum1)                       // arr数组元素的和是:171
	fmt.Printf("平均值：%.2f\n", float64(sum1)/float64(len(arr))) // 平均值：34.20

	//2、请求出一个数组的最大值，并得到对应的下标
	/*
		思路
		1、 声明一个数组 var intArr = [...]int {1, -1, 12, 65, 11}
		2、假定第一个元素就是最大值，下标就 0
		3、然后从第二个元素开始循环比较，如果发现有更大，则交换
	*/

	var intArr = [...]int{1, -1, 12, 65, 11}
	max := intArr[0]
	index := 0
	for i := 1; i < len(intArr); i++ {
		if max < intArr[i] {
			max = intArr[i]
			index = i
		}
	}
	fmt.Printf("最大值是：%v，下标是：%v\n", max, index) // 最大值是：65，下标是：3

	//3、从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	var arr1 = [...]int{1, 3, 5, 7, 8}

	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i]+arr1[j] == 8 {
				fmt.Printf("(%v, %v)", i, j)
			}
		}
	}
}
