package main

import "fmt"

func main() {
	// 请写出下面代码的输出结果
	var a = make([]string, 5, 10)
	for i := 0; i < 12; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("a = %v, len(a) = %v, cap(a) = %v\n", a, len(a), cap(a)) // a = [     0 1 2 3 4 5 6 7 8 9 10 11], len(a) = 17, cap(a) = 20
}
