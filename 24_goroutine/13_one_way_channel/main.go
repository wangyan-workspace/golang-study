package main

import "fmt"

// 单向管道
func main() {
	// 1、在默认情况下下，管道是双向的，可读可写
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 29
	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2) // 10 29

	// 2、管道声明为只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 12
	// <-ch2 // rcannot receive from send-only channel ch2

	// 3、管道声明为只读
	ch3 := make(<-chan int, 2)
	// ch3 <- 10 // cannot send to receive-only channel ch3
	fmt.Println(ch3) // 0x14000110100
}
