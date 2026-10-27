package main

import "fmt"

//函数作为返回值
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

//定义一个方法类型
type calcType func(int, int) int

func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}
 	default:
		return nil
	}
}

func main() {
	var a = do("+")
	fmt.Println(a(10, 20)) // 30

	b := do("*")
	fmt.Println(b(10, 20)) // 200

	c := do("-")
	fmt.Println(c(10, 20)) // -10
}