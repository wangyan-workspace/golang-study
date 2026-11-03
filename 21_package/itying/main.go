package main

import (
	"fmt"
	"itying/calc"
	"itying/tools"
)

func main() {
	//调用calc包里面的方法
	sum := calc.Add(10, 20)
	fmt.Println(sum) // 30
	// fmt.Println(calc.aa) //私有变量不能访问
	fmt.Println(calc.Age) // 18

	sub := calc.Sub(10, 2)
	fmt.Println(sub) // 8

	//调用tools包里面的方法
	b := tools.Mul(10, 2)
	fmt.Println(b) // 20
}
