package main

import "fmt"

// 循环遍历管道数据
func main() {
	// 使用for range遍历通道，当通道被关闭的时候就会退出for range,
	// 如果没有关闭管道就会报个错误fatal error: all goroutines are asleep - deadlock!
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}

	close(ch1)

	// for range循环遍历管道的值  ,注意：管道没有key
	for v := range ch1 {
		fmt.Println(v)
	}

	fmt.Println("--------------------------------")
	var ch2 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch2 <- i
	}
	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)
	}

	/*
		第一次遍历：
		1
		2
		3
		4
		5
		6
		7
		8
		9
		10
		--------------------------------
		第二次遍历：
		1
		2
		3
		4
		5
		6
		7
		8
		9
		10
	*/
}
