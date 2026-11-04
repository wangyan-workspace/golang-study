package main

import "fmt"

type Usber interface {
	start()
	stop()
}

// 电脑结构体
type Computer struct{}

func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

// 手机结构体
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "start")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "stop")
}

// 照相机结构体
type Camera struct{}

func (c Camera) start() {
	fmt.Println("Camera start")
}
func (c Camera) stop() {
	fmt.Println("Camera stop")
}

func main() {
	var computer = Computer{}
	var phone = Phone{
		Name: "iPhone",
	}
	var camera = Camera{}

	computer.work(phone)  // iPhone start, iPhone stop
	computer.work(camera) // Camera start, Camera stop
}
