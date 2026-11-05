package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num // 将素数写入primeChan
		}
	}
	// 如果一个channel关闭了就没法给这个channel发送数据了
	// 什么时候关闭primeChan?
	// 给exitChan里面放入一条数据
	exitChan <- true
	wg.Done()
}

func printPrime(primeChan chan int) {
	// for v := range primeChan {
	// 	fmt.Println(v)
	// }
	wg.Done()
}

func main() {
	start := time.Now().UnixMilli()
	var intChan = make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, 16) // 用于判断协程是否完成 标识primeChan close

	// 1、开启一个协程，向intChan中写入1000个数据
	wg.Add(1)
	go putNum(intChan)

	// 2、开启多个协程，从intChan中取出数据，并判断是否为素数，如果是素数，则写入primeChan
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	// 3、开启一个协程，从primeChan中取出数据，并打印
	wg.Add(1)
	go printPrime(primeChan)

	// 4、判断exitChan是否存满值
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		// 关闭primeChan
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("主线程退出")

	end := time.Now().UnixMilli()
	fmt.Println("执行时间：", end-start, "ms") // 执行时间： 101 ms   95 ms   80 ms  83 ms   81 ms
}
