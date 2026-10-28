package main

import (
	"time"
	"fmt"
)

func main() {
	/*
		golang定时器
	*/
	ticker := time.NewTicker(time.Second)
	
	/*
		每隔一秒打印一次当前时间
		2026-10-27 17:49:51.831525708 +0800 CST m=+1.000216917
		2026-10-27 17:49:52.831527375 +0800 CST m=+2.000215917
		2026-10-27 17:49:53.831532208 +0800 CST m=+3.000217250
		2026-10-27 17:49:54.831534292 +0800 CST m=+4.000216334
		2026-10-27 17:49:55.831539292 +0800 CST m=+5.000218543
		2026-10-27 17:49:56.831540042 +0800 CST m=+6.000216418
		2026-10-27 17:49:57.831540917 +0800 CST m=+7.000214709
		...
	*/
	for t := range ticker.C {
		fmt.Println(t)
	}


	ticker1 := time.NewTicker(time.Second)
	n := 5
	/*
	每隔一秒打印一次当前时间，打印5次后停止
	2026-10-27 18:10:48.044916416 +0800 CST m=+1.000172584
	2026-10-27 18:10:49.044918333 +0800 CST m=+2.000170293
	2026-10-27 18:10:50.044932041 +0800 CST m=+3.000179042
	2026-10-27 18:10:51.044928583 +0800 CST m=+4.000170376
	2026-10-27 18:10:52.044931666 +0800 CST m=+5.000169167
	*/
	for t := range ticker1.C {
		n--
		fmt.Println(t)
		if n == 0 {
			ticker1.Stop() //终止这个定时器继续执行
			break
		}
	}

	// 休眠方法
	fmt.Println("aa")
	time.Sleep(time.Second)
	fmt.Println("bb")
	time.Sleep(time.Second)
	fmt.Println("cc")
	time.Sleep(time.Second * 5)
	fmt.Println("dd")

	for {
		time.Sleep(time.Second)
		fmt.Println("我在定时执行任务")
	}
}