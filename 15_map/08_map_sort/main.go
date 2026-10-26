package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的排序
	var map1 = make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	map1[11] = 100
	map1[2] = 13
	fmt.Println(map1) // map[1:13 2:13 4:56 8:90 10:100 11:100]

	/*
		4 56
		8 90
		2 13
		11 100
		10 100
		1 13
		---------------
		10 100
		1 13
		4 56
		8 90
		11 100
		2 13
	*/
	for k, v := range map1 {
		fmt.Println(k,v) // 多次打印的顺序是不一样的
	}

	//需求：按照key升序输出map的key=>value   (签名算法)

	//1、把map的key放在切片里面
	var keySlice []int
	for key, _ := range map1 {
		keySlice = append(keySlice, key)
	}
	fmt.Println(keySlice) // [11 4 2 10 1 8]

	//2、让key进行升序排序
	sort.Ints(keySlice)
	fmt.Println(keySlice) // [1 2 4 8 10 11]

	//3、循环遍历key输出map的值
	for _, v := range keySlice {
		/*
			key=1 value=13
			key=2 value=13
			key=4 value=56
			key=8 value=90
			key=10 value=100
			key=11 value=100
		*/
		fmt.Printf("key=%v value=%v\n", v, map1[v])
	} 
}