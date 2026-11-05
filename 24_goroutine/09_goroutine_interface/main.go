package main

import (
	"fmt"
	"sync"
)

type Cat struct {
	Name string
	Age  int
}

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 1; i < 10; i++ {
		intChan <- i
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "hello"
	cat := Cat{"小花猫", 4}
	allChan <- cat

	//我们希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat=%v type=%T\n", newCat, newCat) // newCat={小花猫 4} type=main.Cat
	// 下面的写法是错误的!编译不通过
	// fmt.Println("newCat.Name=", newCat.Name) // newCat.Name undefined (type interface{} has no field or method Name)
	// 使用类型断言
	cat1 := newCat.(Cat)
	fmt.Println("cat1.Name=", cat1.Name) // cat1.Name= 小花猫

	var intchan = make(chan int, 1000)
	wg.Add(1)
	go putNum(intchan)
	wg.Wait()
}
