package main

import "fmt"

func main() {
	// Go 语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
	slice := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为 2 的元素  删除的元素：32  注意：append合并切片的时候最后一个元素要加...
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice) // [30 31 33 34 35 36 37]

	//要删除35, 36
	sliceA := []int{30, 31, 32, 33, 34, 35, 36, 37}
	sliceA = append(sliceA[:5], sliceA[7:]...)
	fmt.Println(sliceA) // [30 31 32 33 34 37]

}
