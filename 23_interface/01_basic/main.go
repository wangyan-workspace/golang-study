package main

import "fmt"

// 接口是一个规范
type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

// 手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "start")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "stop")
}

type Camera struct{}

func (c Camera) start() {
	fmt.Println("Camera start")
}
func (c Camera) stop() {
	fmt.Println("Camera stop")
}
func (c Camera) run() {
	fmt.Println("Camera run")
}

func main() {
	p := Phone{
		Name: "iPhone",
	}
	var p1 Usber //golang中接口就是一个数据类型
	p1 = p       //表示手机实现Usb接口
	p1.start()   // iPhone start
	p1.stop()    // iPhone stop

	c := Camera{}
	var c1 Usber
	c1 = c     //表示相机实现Usb接口
	c1.start() // Camera start
	c1.stop()  // Camera stop
	// c1.run()   //错误：type Usber has no field or method run)

	c.run() // Camera run
}
