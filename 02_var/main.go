package main

import "fmt"

var globalVar string = "全局变量" // 全局变量
// gl := 10 // 短变量只能用于声明局部变量，不能用于全局变量的声明

func getUserInfo() (string, int) {
	return "lucky", 20
}

func main() {
	var name string            // 变量声明后没有初始化的话  值为空
	fmt.Println("name=", name) // name=
	name = "lucky"
	fmt.Println("name=", name) // name= lucky

	// var 1a int // 变量名不能以数字开头(错误写法)
	// fmt.Println("1a=", 1a)

	// var if = 10 // 变量名不能以关键字开头(错误写法)
	// fmt.Println("if=", if)

	// go语言变量的定义以及初始化
	// 第一种初始化变量的方式
	var age int
	age = 18
	fmt.Println("age=", age) // age= 18

	// 第二种初始化变量的方式
	var age1 int = 28
	fmt.Println("age1=", age1) // age1= 28

	// 第三种初始化变量的方式
	var age2 = 38
	fmt.Println("age2=", age2) // age2= 38

	// Go 语言中的变量需要声明后才能使用，同一作用域内不支持重复声明
	var username = "lucky"
	var age3 = 20
	var sex = "女"
	fmt.Println(username, age3, sex) // lucky 20 女
	// var username = "zhangsan" // 重复声明会报错

	var a1, a2 string
	a1 = "aaa"
	a2 = "bbb"
	fmt.Println(a1, a2) // aaa bbb

	/*
		var a3, a4 string
		a1 = "ccc"
		a2 = 3 // //错误写法  go语言中定义的类型是string类型 赋值的时候必须赋值string类型的数据
		fmt.Println(a3, a4)
	*/

	var (
		username1 string
		age4      int
		sex1      string
	)

	username1 = "lucky"
	age4 = 20
	sex1 = "女"
	fmt.Println(username1, age4, sex1) // lucky 20 女

	var (
		username2 = "Jack"
		age5      = 30
	)
	fmt.Println(username2, age5) // Jack 30

	/*
		短变量声明法 在函数内部，可以使用更简略的 := 方式声明并初始化变量
		注意：短变量只能用于声明局部变量，不能用于全局变量的声明
	*/

	fmt.Println(globalVar) // 全局变量
	// fmt.Println(gl)        // 短变量只能用于声明局部变量，不能用于全局变量的声明

	username3 := "Tom"
	fmt.Println(username3)        // Tom
	fmt.Printf("类型%T", username3) // 类型string

	// 使用端变量一次声明多个变量，并初始化变量
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c) // 1 2 3

	d, e, f := 1, "hello", 3.14
	fmt.Println(d, e, f)                           // 1 hello 3.14
	fmt.Printf("d的类型:%T e的类型:%T f的类型:%T", d, e, f) // d的类型:int e的类型:string f的类型:float64

	/*
		匿名变量 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）
		匿名变量用一个下划线_表示
	*/

	var username4, age6 = getUserInfo()
	fmt.Println(username4, age6) // lucky 20

	var username5, _ = getUserInfo()
	fmt.Println(username5) // lucky
	//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	_, age7 := getUserInfo()
	fmt.Println(age7) // 20

}
