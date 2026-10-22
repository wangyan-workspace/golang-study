package main

import "fmt"

func main() {
	/*
		for 初始语句;条件表达式;结束语句{
			循环体语句
		}
	*/

	// 1、打印1-10的所有数据

	/*
		for循环的执行流程
			1、i := 1
			2、i <= 10
			3、执行花括号里面的语句
			4、i++
			5、i <= 10
			6、执行花括号里面的语句
			7、i++
			8、i <= 10
	*/

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2、for 循环的初始语句可以被忽略，但是初始语句后的分号必须要写
	i := 1
	for ; i <= 10; i++ {
		fmt.Println(i)
	}

	/*
		写for循环的时候要注意死循环

		for 条件 {
			循环体语句
		}

		for 循环的初始语句和结束语句都可以省略
	*/
	j := 1
	for j <= 10 {
		fmt.Printf("j=%v\n", j)
		j++
	}

	/*
		for {
			循环体语句
		}
		注意：Go 语言中是没有 while 语句的，我们可以通过 for 代替
	*/
	k := 1
	for {
		if k <= 10 {
			fmt.Printf("k=%v\n", k)
		} else {
			break // break 语句用于跳出循环
		}
		k++
	}
}
