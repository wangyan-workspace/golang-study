package main

import "fmt"

func main() {
	/*

		switch expression {
		case condition:
				匹配执行语句
				break
		case condition:
				匹配执行语句
				break
		case condition:
				匹配执行语句
				break
		default:
				break
		}
	*/

	/*
		练习：判断文件类型
		如果后缀名是.html 输入 text/html
		如果后缀名.css 输出 text/css
		如果 后缀名是.js 输出 text/javascript
	*/
	var ext = ".css"
	switch ext {
	case ".html":
		fmt.Println("text/html")
		// break // break可以省略
	case ".css":
		fmt.Println("text/css")
		// break
	case ".js":
		fmt.Println("text/javascript")
		// break
	default:
		fmt.Println("unknown")
		// break
	}

	// 2、switch case的另一种写法
	switch extname := ".js"; extname {
	case ".html":
		fmt.Println("text/html")
		// break // break可以省略
	case ".css":
		fmt.Println("text/css")
		// break
	case ".js":
		fmt.Println("text/javascript")
		// break
	default:
		fmt.Println("unknown")
		// break
	}
	// fmt.Println(extname) //undefined: extname

	// 3、一个分支可以有多个值，多个 case 值中间使用英文逗号分隔
	// 判断一个数是不是偶数
	var n = 8
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("未知")
	}

	var score = "B"
	switch score {
	case "A", "B", "C":
		fmt.Println("及格")
	case "D":
		fmt.Println("不及格")
	}

	//ABC及格  D不及格
	switch score := "C"; score {
	case "A", "B", "C":
		fmt.Println("及格")
	case "D":
		fmt.Println("不及格")
	}

	//4、分支还可以使用表达式，这时候 switch 语句后面不需要再跟判断变量
	var age = 80 // 注意身体
	switch {
	case age < 18:
		fmt.Println("好好成长")
	case age >= 18 && age < 24:
		fmt.Println("好好学习")
	case age >= 24 && age <= 60:
		fmt.Println("好好工作")
	case age > 60 && age < 80:
		fmt.Println("好好享受生活")
	default:
		fmt.Println("注意身体")
	}

	//5、 switch 的穿透 fallthrought
	//fallthrough`语法可以执行满足条件的 case 的下一个 case，是为了兼容 C 语言中的 case 设计 的
	/*
		好好工作
		好好享受生活
	*/
	var age1 = 40
	switch {
	case age1 < 18:
		fmt.Println("好好成长")
	case age1 >= 18 && age1 < 24:
		fmt.Println("好好学习")
	case age1 >= 24 && age1 <= 60:
		fmt.Println("好好工作")
		fallthrough
	case age1 > 60 && age1 < 80:
		fmt.Println("好好享受生活")
	default:
		fmt.Println("注意身体")
	}
}
