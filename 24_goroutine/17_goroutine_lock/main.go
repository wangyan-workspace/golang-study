package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0

var wg sync.WaitGroup

var mutex sync.Mutex

func test() {
	mutex.Lock()
	count++
	fmt.Println("the count is : ", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()

	/*
		并发安全问题，容易导致结果不准确，加锁可以解决这个问题,以下是加锁后的结果
		the count is :  1
		the count is :  2
		the count is :  3
		the count is :  4
		the count is :  5
		the count is :  6
		the count is :  7
		the count is :  8
		the count is :  9
		the count is :  10
		the count is :  11
		the count is :  12
		the count is :  13
		the count is :  14
		the count is :  15
		the count is :  16
		the count is :  17
		the count is :  18
		the count is :  19
		the count is :  20
	*/
}
