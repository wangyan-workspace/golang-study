package main

import "fmt"

//一个结构体中可以嵌套包含另一个结构体或结构体指针

//父亲结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%v在跑\n", a.Name)
}

//子结构体
type Dog struct {
	Age int
	*Animal //结构体嵌套 继承
}

func (d Dog) bark() {
	fmt.Printf("%v在叫\n", d.Name)
}

func main() {
	var dog = Dog{
		Age: 1,
		Animal: &Animal{
			Name: "旺财",
		},
	}
	dog.run() // 旺财在跑
	dog.bark() // 旺财在叫
}