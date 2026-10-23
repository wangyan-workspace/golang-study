package main

import "fmt"

func main() {
	var slice = make([]int, 4, 8)
	fmt.Println(slice) // [0 0 0 0]
	slice[0] = 10
	slice[1] = 12
	slice[2] = 89
	slice[3] = 30
	fmt.Println(slice) // [10 12 89 30]

	sliceA := []string{"one", "two", "three"}
	sliceA[2] = "four"
	fmt.Println(sliceA) // [one two four]

	//golang中没法通过下标的方式给切片扩容
	var sliceB []int
	fmt.Printf("%v---%v---%v\n", sliceB, len(sliceB), cap(sliceB)) // []---0---0
	// sliceB[0] = 1 // 错误写法
	fmt.Println(sliceB) // []
}
