package main

import "fmt"

func main() {
	//1、基于数组定义切片
	var arr = [5]int{1, 2, 3, 4, 5}
	var slice = arr[:]                  //[:]表示取全部元素 获取数组里面的所有值
	fmt.Printf("%v-%T\n", slice, slice) // [1 2 3 4 5]-[]int

	slice1 := arr[1:4]                    // 取下标1到下标3的元素,是一个左闭右开的区间
	fmt.Printf("%v-%T\n", slice1, slice1) // [2 3 4]-[]int

	slice2 := arr[2:]                     // 取下标2到最后的元素
	fmt.Printf("%v-%T\n", slice2, slice2) // [3 4 5]-[]int

	slice3 := arr[:3]                     // 取下标0到下标3的元素
	fmt.Printf("%v-%T\n", slice3, slice3) // [1 2 3]-[]int

	//2、基于切片定义切片
	slice4 := []string{"hello", "world", "golang"}
	slice5 := slice4[1:]
	fmt.Printf("%v-%T\n", slice5, slice5) // [world golang]-[]string

	//3、关于切片的长度和容量
	// 长度：切片的长度就是它所包含的元素个数
	// 容量：切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

	s1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度%d 容量%d\n", len(s1), cap(s1)) // 长度6 容量6

	s2 := s1[2:]                                // 5, 7, 11, 13
	fmt.Printf("长度%d 容量%d\n", len(s2), cap(s2)) // 长度4 容量4

	s3 := s1[1:3]                               // 3, 5
	fmt.Printf("长度%d 容量%d\n", len(s3), cap(s3)) // 长度2 容量5

	s4 := s1[:3]                                // 2, 3, 5
	fmt.Printf("长度%d 容量%d\n", len(s4), cap(s4)) // 长度3 容量6
}
