package main

import "fmt"

/*
结构体指针接收者实现接口：

指针接收者： 如果结构体中的方法是指针接收者，那么实例化后结构体指针类型都可以赋值给接口变量， 结构体值类型没法赋值给接口变量。

*/

type Usber interface {
	start()
	stop()
}

// 手机结构体
type Phone struct {
	Name string
}

func (p *Phone) start() { // 指针接收者
	fmt.Println(p.Name, "start")
}
func (p *Phone) stop() { // 指针接收者
	fmt.Println(p.Name, "stop")
}

func main() {
	/*
		错误写法
		var phone = Phone{
			Name: "iPhone",
		}
		var p1 Usber = phone // Phone does not implement Usber (method start has pointer receiver)
		p1.start()
	*/
	var phone1 = &Phone{
		Name: "iPhone",
	}
	var p1 Usber = phone1
	p1.start() // iPhone start
}
