package main

import "fmt"

func main() {
	//make()函数创建一个切片  make([]T, size, cap)
	var slice = make([]int, 4, 8)
	fmt.Println(slice)                            // [0 0 0 0]
	fmt.Printf("%d---%d", len(slice), cap(slice)) // 4---8
}
