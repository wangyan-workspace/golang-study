package main

import "fmt"

func main() {
	//切片的循环遍历
	var slice = []string{"php", "go", "python", "java"}

	/*
		php
		go
		python
		java
	*/
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	/*
		0 php
		1 go
		2 python
		3 java
	*/
	for k, v := range slice {
		fmt.Println(k, v)
	}
}
