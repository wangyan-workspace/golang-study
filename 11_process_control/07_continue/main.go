package main

import "fmt"

func main() {
	// continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用
	/*
		1
		2
		4
		5
		6
		7
		8
		9
		10
	*/
	for i := 1; i <= 10; i++ {

		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	/*
		i=0 j=0
		i=0 j=1
		i=0 j=2
		i=0 j=4
		i=0 j=5
		i=0 j=6
		i=0 j=7
		i=0 j=8
		i=0 j=9
		i=1 j=0
		i=1 j=1
		i=1 j=2
		i=1 j=4
		i=1 j=5
		i=1 j=6
		i=1 j=7
		i=1 j=8
		i=1 j=9
	*/
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}

	// 在 continue 语句后添加标签时，表示开始标签对应的循环

	/*
		i=0 j=0
		i=0 j=1
		i=0 j=2
		i=1 j=0
		i=1 j=1
		i=1 j=2
	*/
label2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue label2
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}
}
