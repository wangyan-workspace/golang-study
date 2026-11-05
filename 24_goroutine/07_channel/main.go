package main

import "fmt"

func main() {
	// 1、创建channel
	ch := make(chan int, 3)

	// 2、给管道里面存储数据
	ch <- 10
	ch <- 20
	ch <- 30

	// 3、获取管道里面的内容
	a := <-ch
	fmt.Println("a=", a) // a= 10

	<-ch // 从管道里面取值 20
	b := <-ch
	fmt.Println("b=", b) // b= 30
	ch <- 56
	ch <- 89

	// 4、打印管道的长度和容量
	fmt.Printf("值：%v 容量：%v 长度%v\n", ch, cap(ch), len(ch)) // 值：0x14000090080 容量：3 长度2

	// 5、管道的类型（引用数据类型）
	ch1 := make(chan int, 4)

	ch1 <- 45
	ch1 <- 34
	ch1 <- 78

	ch2 := ch1

	ch2 <- 100

	<-ch1
	<-ch1
	<-ch1

	c := <-ch1
	fmt.Println("c=", c) // c= 100

	// 6、管道阻塞
	ch3 := make(chan int, 1)
	ch3 <- 98
	// ch3 <- 45 // 超过管道的存储容量了 all goroutines are asleep - deadlock!

	// 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	ch4 := make(chan string, 2)
	ch4 <- "hello"
	ch4 <- "world"

	d := <-ch4
	e := <-ch4
	fmt.Println(d, e) // hello world

	// f := <-ch4
	// fmt.Println("f=", f) // all goroutines are asleep - deadlock!

	// 正确的写法
	ch5 := make(chan int, 1)
	ch5 <- 57
	<-ch5
	ch5 <- 78
	<-ch5
	ch5 <- 23
	g := <-ch5
	fmt.Println("g=", g) // g= 23
}
