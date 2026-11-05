package main

import (
	"fmt"
	"sync"
	"time"
)

//需求：要统计1-120000的数字中那些是素数？goroutine  for循环实现

/*
1 协程  统计  1-30000

2 协程  统计  30001-60000

3 协程  统计  60001-90000

4 协程  统计  90001-120000

// start:(n-1)*30000+1       end:n*30000
*/

var wg sync.WaitGroup

func test(n int) {
	for num := (n-1)*30000 + 1; num <= n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				// fmt.Println(num, "是素数")
			}
		}
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixMilli()

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test(i)
	}

	wg.Wait()

	fmt.Println("关闭主线程...")
	end := time.Now().UnixMilli()
	fmt.Println("程序执行时间：", end-start, "毫秒") // 程序执行时间： 179 毫秒   185 毫秒  159 毫秒  159 毫秒
}
