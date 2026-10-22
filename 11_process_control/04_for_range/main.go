package main

import "fmt"

func main() {
	var str = "你好golang"

	/*
		key:0, value:你
		key:3, value:好
		key:6, value:g
		key:7, value:o
		key:8, value:l
		key:9, value:a
		key:10, value:n
		key:11, value:g
	*/
	for i, v := range str {
		fmt.Printf("key:%d, value:%c\n", i, v)
	}

	/*
		hello
		world
		golang
	*/
	var arr = []string{"hello", "world", "golang"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	/*
		php
		java
		node
		golang
	*/
	var arr1 = []string{"php", "java", "node", "golang"}
	for _, val := range arr1 {
		fmt.Println(val)
	}
}
