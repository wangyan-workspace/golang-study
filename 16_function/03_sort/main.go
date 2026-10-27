package main

import "fmt"

func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if(slice[i] > slice[j]) {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}

	return slice
}

func sortIntDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if(slice[i] < slice[j]) {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}

	return slice
}

func main() {
	//案例1：把前面讲的选择排序封装成方法，实现整型切片的升序、降序排序排列

	var sliceA = []int{12, 34, 37, 35, 556, 36, 2}
	arrA := sortIntAsc(sliceA)
	fmt.Println(arrA) // [2 12 34 35 36 37 556]

	var sliceB = []int{1, 34, 4, 36, 126, 36, 2}
	sortIntDesc(sliceB)
	fmt.Println(sliceB) // [126 36 36 34 4 2 1]
}