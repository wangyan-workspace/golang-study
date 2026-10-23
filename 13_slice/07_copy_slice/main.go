package main

import "fmt"

func main() {
	//1、copy()函数复制切片
	sliceA := []int{1, 2, 3, 4, 5, 6}
	sliceB := make([]int, 4, 4)
	copy(sliceB, sliceA)
	fmt.Println(sliceA) // [1 2 3 4 5 6]
	fmt.Println(sliceB) // [1 2 3 4]
	sliceB[0] = 111
	fmt.Println(sliceA) // [1 2 3 4 5 6]
	fmt.Println(sliceB) // [111 2 3 4]
}
