package main

import "fmt"

func main() {
	// 切片的声明 初始化
	var s []int
	fmt.Println(s)                                // []
	fmt.Printf("%v - %T - 长度:%v\n", s, s, len(s)) // [] - []int - 长度:0

	var s1 = []int{12, 34, 46, 67}
	fmt.Printf("%v - %T - 长度:%v\n", s1, s1, len(s1)) //  [12 34 46 67] - []int - 长度:4

	var s3 = []int{1: 2, 2: 4, 5: 6}
	fmt.Printf("%v - %T - 长度:%v\n", s3, s3, len(s3)) //  [0 2 4 0 0 6] - []int - 长度:6

	var s4 []int
	var s5 = []int{1, 2, 3, 4}
	fmt.Println(s4)        // []
	fmt.Println(s4 == nil) // true golnag中申明切片以后 切片的默认值就是nil
	fmt.Println(s5 == nil) // false
}
