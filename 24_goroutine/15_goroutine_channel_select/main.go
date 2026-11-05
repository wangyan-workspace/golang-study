package main

import (
	"fmt"
	"time"
)

func main() {
	// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复用

	// 1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	fmt.Println("intChan:", intChan)       // intChan: 0x1400013a000
	fmt.Println("stringChan:", stringChan) // stringChan: 0x1400013c000

	//使用select来获取channel里面的数据的时候不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan中读取的数据%d\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Printf("从stringChan中读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Printf("数据获取完毕")
			return //注意退出...
		}
	}

	/*
		从intChan中读取的数据0
		从intChan中读取的数据1
		从intChan中读取的数据2
		从intChan中读取的数据3
		从stringChan中读取的数据hello0
		从intChan中读取的数据4
		从stringChan中读取的数据hello1
		从intChan中读取的数据5
		从stringChan中读取的数据hello2
		从stringChan中读取的数据hello3
		从intChan中读取的数据6
		从stringChan中读取的数据hello4
		从intChan中读取的数据7
		从intChan中读取的数据8
		从intChan中读取的数据9
		数据获取完毕
	*/
}
