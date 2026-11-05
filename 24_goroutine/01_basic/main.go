package main

import (
	"fmt"
	"time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔50毫秒秒输出 "你好golang"
// 在主线程中也每隔50毫输出"你好golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test()你好golang", i)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	go test() //表示开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("main()你好golang", i)
		time.Sleep(time.Millisecond * 50)
	}

	// 从结果上来看，test() 和 main() 是交替执行的，并不是main() 执行完后，再执行test()
	// 因为test() 和 main() 是交替执行的，所以test() 和 main() 的执行顺序是不确定的
	// 这就是goroutine的特性，goroutine是并发执行的，而不是顺序执行的

	/*
		程序执行结果打印：
			main()你好golang 0
			test()你好golang 0
			test()你好golang 1
			main()你好golang 1
			test()你好golang 2
			main()你好golang 2
			test()你好golang 3
			main()你好golang 3
			test()你好golang 4
			main()你好golang 4
			main()你好golang 5
			test()你好golang 5
			test()你好golang 6
			main()你好golang 6
			main()你好golang 7
			test()你好golang 7
			test()你好golang 8
			main()你好golang 8
			main()你好golang 9
			test()你好golang 9
	*/

	/*
		程序执行结果打印：
			main()你好golang 0
			test()你好golang 0
			test()你好golang 1
			main()你好golang 1
			main()你好golang 2
			test()你好golang 2
			test()你好golang 3
			main()你好golang 3
			main()你好golang 4
			test()你好golang 4
			test()你好golang 5
			main()你好golang 5
			main()你好golang 6
			test()你好golang 6
			test()你好golang 7
			main()你好golang 7
			main()你好golang 8
			test()你好golang 8
			test()你好golang 9
			main()你好golang 9
	*/
}
