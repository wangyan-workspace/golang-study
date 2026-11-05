package main

import (
	"fmt"
	"sync"
	"time"
)

// 主线程退出后所有的协程无论有没有执行完毕都会退出，
// 所以我们在主进程中可以通过WaitGroup等待协程执行完毕
var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1()你好golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() //协程计数器-1
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2()你好golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() //协程计数器-1
}

func main() {
	wg.Add(1)  //协程计数器+1
	go test1() //表示开启一个协程

	wg.Add(1)  //协程计数器+1
	go test2() //表示开启一个协程

	wg.Wait() //等待协程执行完毕

	fmt.Println("主线程退出...")

	/*
		程序执行结果：
			test2()你好golang- 0
			test1()你好golang- 0
			test1()你好golang- 1
			test2()你好golang- 1
			test2()你好golang- 2
			test1()你好golang- 2
			test1()你好golang- 3
			test2()你好golang- 3
			test2()你好golang- 4
			test1()你好golang- 4
			test1()你好golang- 5
			test2()你好golang- 5
			test2()你好golang- 6
			test1()你好golang- 6
			test1()你好golang- 7
			test2()你好golang- 7
			test2()你好golang- 8
			test1()你好golang- 8
			test1()你好golang- 9
			test2()你好golang- 9
			主线程退出...
	*/
}
