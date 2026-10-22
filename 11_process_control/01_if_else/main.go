package main

import "fmt"

func main() {
	/*
		Go 语言中 if 条件判断的格式如下：
		if 表达式1 {
			分支 1
		}else if 表达式2 {
			分支 2
		}else{
			分支 3
		}
	*/

	//1、最简单的if语句
	flag := true
	if flag {
		fmt.Println("flag is true")
	}

	age := 20
	if age > 18 {
		fmt.Println("成年人")
	}

	//2、if语句的另一种写法
	if num := 30; num > 18 {
		fmt.Println("成年人")
	}

	// 3、探讨上面两种写法的区别
	/*
		age := 20  //当前区域内是全局变量
		if age > 18 {
			fmt.Println("成年人")
		}
	*/
	/*
		if num := 30; num > 18 { //局部变量
			fmt.Println("成年人")
		}
	*/

	//4、输入一个人的成绩，如果成绩大于等于90输出A，如果小于90大于75输出B,否则输出C
	var score int = 85
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	if score = 60; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// if else要注意的细节
	/*
		1、if{}不能省略掉，即使只有一行代码也要加上{}
		2、{ 必须紧挨着条件
	*/

	/*
		age1 := 30
		if age1 > 18  // 错误写法
			fmt.Println("成年人")
	*/

	//6、求两个数的最大值
	var a, b int = 10, 20
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Println("max=", max)
}
