package main

import "fmt"

// 接口嵌套
type Ainterface interface {
	SetName(string)
}
type Binterface interface {
	GetName() string
}
type Animaler interface {
	Ainterface
	Binterface
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

func main() {
	//Dog实现Animal的接口
	d := &Dog{
		Name: "旺财",
	}
	var d1 Animaler = d //表示让Dog实现Animaler这个接口
	d1.SetName("大黄")
	fmt.Println(d1.GetName()) //大黄
}
