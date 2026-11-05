package main

import (
	"fmt"
	"sync"
	"time"
)

/*
需求：使用goroutine和channel协同工作案例
1、开启一个WriteData的的协程给向管道inChan中写入100条数据
2、开启一个ReadData的协程读取inChan中写入的数据
3、注意：WriteData和ReadData同时操作一个管道
4、主线程必须等待操作完成后才可以退出

goroutine结合Channel使用的简单demo,定义两个方法，一个方法给管道里面写数据，一个给管道里面读取数据。要求同步进行。
*/

var wg sync.WaitGroup

// 写数据
func writeData(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("【写入】数据%v成功\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
	wg.Done()
}

// 读数据
func readData(ch chan int) {
	for v := range ch {
		fmt.Printf("【读取】数据%v成功\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func main() {
	var ch = make(chan int, 10)

	wg.Add(1)
	go writeData(ch)

	wg.Add(1)
	go readData(ch)

	wg.Wait()
	fmt.Println("主线程退出")

	/*
		结果：
		【写入】数据1成功
		【读取】数据1成功
		【写入】数据2成功
		【读取】数据2成功
		【写入】数据3成功
		【读取】数据3成功
		【写入】数据4成功
		【读取】数据4成功
		【写入】数据5成功
		【读取】数据5成功
		【写入】数据6成功
		【读取】数据6成功
		【写入】数据7成功
		【读取】数据7成功
		【读取】数据8成功
		【写入】数据8成功
		【写入】数据9成功
		【读取】数据9成功
		【写入】数据10成功
		【读取】数据10成功
		主线程退出
	*/
}
