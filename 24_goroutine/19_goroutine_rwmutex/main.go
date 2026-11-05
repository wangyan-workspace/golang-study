package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

func write() {
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

func read() {
	mutex.RLock()
	fmt.Println("---执行读操作")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
	wg.Done()
}

func main() {
	//开启10个协程执行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	// //开启10个协程执行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()

	// 写操作是串行执行的，读操作是并行执行的
	/*
		执行写操作
		---执行读操作
		---执行读操作
		---执行读操作
		---执行读操作
		---执行读操作
		---执行读操作
		---执行读操作
		---执行读操作
		---执行读操作
		---执行读操作
		执行写操作
		执行写操作
		执行写操作
		执行写操作
		执行写操作
		执行写操作
		执行写操作
		执行写操作
		执行写操作
	*/
}
